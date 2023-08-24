//go:build linux

package extruntime

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func logCapsInformation(level zerolog.Level) {
	caps, err := exec.Command("getpcaps", strconv.Itoa(os.Getpid())).CombinedOutput()
	if err == nil {
		log.WithLevel(level).Msgf("Process capabilities: %s", string(caps))
	} else {
		log.WithLevel(level).Msgf("Process capabilities: %s", err.Error())
	}
}

func logUnameInformation(level zerolog.Level) {
	var uts syscall.Utsname
	syscall.Uname(&uts)

	log.WithLevel(level).Msgf(
		"OS information: sysname=%s; release=%s; version=%s; machine=%s",
		charsToString(uts.Sysname[:]),
		charsToString(uts.Release[:]),
		charsToString(uts.Version[:]),
		charsToString(uts.Machine[:]),
	)
}

func charsToString(ca []int8) string {
	s := make([]byte, len(ca))
	i := 0
	for ; i < len(ca); i++ {
		if ca[i] == 0 {
			break
		}
		s[i] = uint8(ca[i])
	}
	return string(s[:i])
}
