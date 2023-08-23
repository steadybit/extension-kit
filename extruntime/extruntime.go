package extruntime

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"runtime"
)

func LogRuntimeInformation(level zerolog.Level) {
	if !log.WithLevel(level).Enabled() {
		return
	}

	log.WithLevel(level).Msgf("Go Runtime information: os=%s; arch=%s", runtime.GOOS, runtime.GOARCH)
	log.WithLevel(level).Msgf("Process information: pid=%d; uid=%d; gid=%d", os.Getpid(), os.Getuid(), os.Getgid())

	logUnameInformation(level)
	logCapsInformation(level)
}
