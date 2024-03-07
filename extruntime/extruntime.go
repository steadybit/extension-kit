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

	log.WithLevel(level).Str("os", runtime.GOOS).Str("arch", runtime.GOOS).Int("GOMAXPROCS", runtime.GOMAXPROCS(0)).Msgf("Go Runtime information")
	log.WithLevel(level).Int("pid", os.Getpid()).Int("uid", os.Getuid()).Int("gid", os.Getgid()).Msgf("Process information")

	logUnameInformation(level)
	logCapsInformation(level)
}

func GetUnameInformation() string {
	return UnameInformation()
}
