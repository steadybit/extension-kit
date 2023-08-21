//go:build linux

package extruntime

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"strconv"
)

func logCapsInformation(level zerolog.Level) {
	caps, err := exec.Command("getpcaps", strconv.Itoa(os.Getpid())).CombinedOutput()
	if err != nil {
		log.WithLevel(level).Msgf("Process capabilities: %s", string(caps))
	} else {
		log.WithLevel(level).Msgf("Process capabilities: %s", err)
	}
}
