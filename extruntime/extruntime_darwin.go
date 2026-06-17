//go:build darwin

package extruntime

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func adjustOOMScoreAdj(_ int) {
	log.Debug().Msg("oom_score_adj adjustment is not supported on this platform")
}

func logCapsInformation(_ zerolog.Level) {
}

func logUnameInformation(_ zerolog.Level) {
}

func UnameInformation() string {
	return ""
}
