package systemd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// RunSystemctl runs `systemctl` with the given args and returns combined output.
func RunSystemctl(args ...string) ([]byte, error) {
	cmd := exec.Command("systemctl", args...)
	out, err := cmd.CombinedOutput()
	return out, err
}

// StreamJournal streams `journalctl` output to stdout/stderr.
func StreamJournal(unit string, follow bool, lines int) error {
	args := []string{"-u", unit}
	if lines > 0 {
		args = append(args, "-n", strconv.Itoa(lines))
	}
	if follow {
		args = append(args, "-f")
	}
	cmd := exec.Command("journalctl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// UnitFromArgs returns the unit name from positional args or from the provided flag value.
func UnitFromArgs(args []string, unitFlag string) (string, error) {
	if len(args) > 0 {
		return args[0], nil
	}
	if unitFlag != "" {
		return unitFlag, nil
	}
	return "", fmt.Errorf("unit name required")
}
