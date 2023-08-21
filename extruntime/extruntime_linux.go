//go:build linux

package extruntime

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/seccomp/libseccomp-golang"
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
func logSeccompInformation(level zerolog.Level) {
	ctx, err := seccomp.NewContext(seccomp.ActLog)
	if err != nil {
		log.WithLevel(level).Msgf("Seccomp: %s", err)
		return
	}
	defer ctx.Release()

	if err := ctx.Load(); err != nil {
		log.WithLevel(level).Msgf("Seccomp: %s", err)
		return
	}

	mode, err := ctx.GetMode()
	if err != nil {
		log.WithLevel(level).Msgf("Seccomp: %s", err)
		return
	}

	log.WithLevel(level).Msgf("Seccomp: mode=%s", err)
}
