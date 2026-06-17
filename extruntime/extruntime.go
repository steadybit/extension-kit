package extruntime

import (
	"context"
	"github.com/elastic/go-sysinfo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"runtime"
	"strconv"
)

const (
	envOOMScoreAdj     = "STEADYBIT_EXTENSION_OOM_SCORE_ADJ"
	defaultOOMScoreAdj = -998
	minOOMScoreAdj     = -1000
	maxOOMScoreAdj     = 1000
)

func LogRuntimeInformation(level zerolog.Level) {
	if !log.WithLevel(level).Enabled() {
		return
	}

	log.WithLevel(level).Str("os", runtime.GOOS).Str("arch", runtime.GOOS).Int("GOMAXPROCS", runtime.GOMAXPROCS(0)).Msgf("Go Runtime information")
	log.WithLevel(level).Int("pid", os.Getpid()).Int("uid", os.Getuid()).Int("gid", os.Getgid()).Msgf("Process information")

	logUnameInformation(level)
	logCapsInformation(level)
}

// AdjustOOMScoreAdj lowers the process oom_score_adj to protect the extension from the Linux
// OOM-killer under node memory pressure. The target score is read from
// STEADYBIT_EXTENSION_OOM_SCORE_ADJ (default -998) and clamped to [-1000, 1000].
// Lowering the score requires CAP_SYS_RESOURCE; without it the adjustment fails and a warning
// is logged. This is a no-op on non-Linux platforms.
func AdjustOOMScoreAdj() {
	adjustOOMScoreAdj(resolveOOMScoreAdj())
}

func resolveOOMScoreAdj() int {
	value, ok := os.LookupEnv(envOOMScoreAdj)
	if !ok {
		return defaultOOMScoreAdj
	}

	score, err := strconv.Atoi(value)
	if err != nil {
		log.Warn().Str(envOOMScoreAdj, value).Msgf("Invalid %s, falling back to %d", envOOMScoreAdj, defaultOOMScoreAdj)
		return defaultOOMScoreAdj
	}

	if score < minOOMScoreAdj {
		return minOOMScoreAdj
	}
	if score > maxOOMScoreAdj {
		return maxOOMScoreAdj
	}
	return score
}

func GetUnameInformation() string {
	return UnameInformation()
}

func GetHostname() (hostname, fqdn string, err error) {
	hostname = "unknown"
	fqdn = "unknown"

	hostname, err = os.Hostname()
	if host, err2 := sysinfo.Host(); err2 == nil {
		fqdn, err = host.FQDNWithContext(context.Background())
	}
	return
}
