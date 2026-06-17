//go:build linux

package extruntime

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

const oomScoreAdjPath = "/proc/self/oom_score_adj"

func adjustOOMScoreAdj(score int) {
	previous := strings.TrimSpace(readOOMScoreAdj())

	if err := os.WriteFile(oomScoreAdjPath, []byte(strconv.Itoa(score)+"\n"), 0644); err != nil {
		log.Warn().Err(err).Int("oom_score_adj", score).Msg("Failed to adjust oom_score_adj; extension may be killed by the OOM killer under memory pressure (missing CAP_SYS_RESOURCE?)")
		return
	}

	log.Info().Int("oom_score_adj", score).Str("previous", previous).Msg("Adjusted oom_score_adj to protect extension from the OOM killer")
}

func readOOMScoreAdj() string {
	if content, err := os.ReadFile(oomScoreAdjPath); err == nil {
		return string(content)
	}
	return "unknown"
}

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
