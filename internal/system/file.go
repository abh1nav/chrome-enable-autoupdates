package system

import (
	"os"
)

// Exists tells you whether a filepath exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err != nil
}
