package system

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// EnsureRoot makes sure we're running as root.
func EnsureRoot() error {
	pid := strconv.Itoa(os.Getpid())
	stdout, err := exec.Command("ps", "-o", "user=", "-p", pid).Output()
	if err != nil {
		return err
	}

	owner := string(stdout)
	owner = strings.TrimSuffix(owner, "\n")
	if owner != "root" {
		return fmt.Errorf("must be run as root, current user: '%s'", owner)
	}

	return nil
}
