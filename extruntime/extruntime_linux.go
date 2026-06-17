//go:build linux

package extruntime

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sys/unix"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

const oomScoreAdjPath = "/proc/self/oom_score_adj"

func adjustOOMScoreAdj(score int) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	raiseSysResourceCapability()

	previous := strings.TrimSpace(readOOMScoreAdj())

	if err := os.WriteFile(oomScoreAdjPath, []byte(strconv.Itoa(score)+"\n"), 0644); err != nil {
		log.Warn().Err(err).Int("oom_score_adj", score).Msg("Failed to adjust oom_score_adj; extension may be killed by the OOM killer under memory pressure (missing CAP_SYS_RESOURCE?)")
		return
	}

	log.Info().Int("oom_score_adj", score).Str("previous", previous).Msg("Adjusted oom_score_adj to protect extension from the OOM killer")
}

// raiseSysResourceCapability best-effort moves CAP_SYS_RESOURCE from the permitted to the
// effective set of the current OS thread. This lets a non-root extension that carries the
// capability as a file capability (setcap cap_sys_resource=+p) lower its oom_score_adj without
// running as root. The caller must lock the OS thread, since capabilities are per-thread and the
// write to oom_score_adj has to happen on the same thread. If the capability is not in the
// permitted set, this is a no-op and the subsequent write fails gracefully with a warning.
func raiseSysResourceCapability() {
	hdr := unix.CapUserHeader{Version: unix.LINUX_CAPABILITY_VERSION_3}
	data := [2]unix.CapUserData{}
	if err := unix.Capget(&hdr, &data[0]); err != nil {
		return
	}

	bit := uint32(1) << uint(unix.CAP_SYS_RESOURCE)
	if data[0].Permitted&bit == 0 {
		return
	}

	data[0].Effective |= bit
	_ = unix.Capset(&hdr, &data[0])
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
