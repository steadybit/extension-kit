package extlogging

import (
	"io"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

var captureLock sync.Mutex

func TestInitZeroLog_Format(t *testing.T) {
	zerolog.TimestampFunc = func() time.Time {
		return time.Date(2020, 1, 1, 0, 0, 59, 99100001, time.UTC)
	}
	defer func() { zerolog.TimestampFunc = time.Now }()

	tests := []struct {
		format       string
		wantedOutput string
	}{
		{format: "json", wantedOutput: "{\"level\":\"info\",\"time\":\"2020-01-01T00:00:59.099Z\",\"message\":\"test\"}\n"},
		{format: "text", wantedOutput: "2020-01-01T00:00:59.099Z INF test\n"},
	}
	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			t.Setenv("STEADYBIT_LOG_FORMAT", tt.format)

			msg, err := captureStdErr(func() {
				InitZeroLog()
				log.Info().Msg("test")
			})

			assert.Nil(t, err)
			assert.Equal(t, tt.wantedOutput, msg)
		})
	}
}

func TestInitZeroLog_Level(t *testing.T) {
	tests := []struct {
		level            string
		wantedLevel      zerolog.Level
		wantedCallerHook bool
	}{
		{level: "debug", wantedLevel: zerolog.DebugLevel, wantedCallerHook: true},
		{level: "", wantedLevel: zerolog.InfoLevel},
	}
	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			t.Setenv("STEADYBIT_LOG_LEVEL", tt.level)

			InitZeroLog()

			assert.Equal(t, tt.wantedLevel, zerolog.GlobalLevel())
		})
	}
}

func captureStdErr(f func()) (string, error) {
	captureLock.Lock()
	defer captureLock.Unlock()

	rescueStderr := os.Stderr
	rescueStdout := os.Stdout
	r, w, err := os.Pipe()

	if err != nil {
		log.Error().Msgf("unable to create os pipe: %s", err)
		return "", err
	}

	os.Stderr = w
	os.Stdout = w

	f()

	defer func() {
		os.Stderr = rescueStderr
		os.Stdout = rescueStdout
	}()

	_ = r.SetReadDeadline(time.Now().Add(100 * time.Millisecond))

	w.Close()
	captured, _ := io.ReadAll(r)
	return string(captured), nil
}
