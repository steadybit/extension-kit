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
	if caps, err := exec.Command("getpcaps", strconv.Itoa(os.Getpid())).CombinedOutput(); err == nil {
		log.WithLevel(level).Msgf("Process capabilities: %s", string(caps))
	}
}

func logUnameInformation(level zerolog.Level) {
	log.WithLevel(level).Msg(UnameInformation())
}

func UnameInformation() string {
	var uts syscall.Utsname
	syscall.Uname(&uts)

	return "OS information: sysname=" + charsToString(uts.Sysname[:]) + "; release=" + charsToString(uts.Release[:]) + "; version=" + charsToString(uts.Version[:]) + "; machine=" + charsToString(uts.Machine[:])
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
