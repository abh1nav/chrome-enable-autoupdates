package chrome

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abh1nav/chrome-enable-autoupdates/internal"
	"github.com/abh1nav/chrome-enable-autoupdates/internal/plist"
)

// InstalledVersion reads the Info.plist and returns the value for CFBundleShortVersionString
// i.e. the installed version of Chrome.
func InstalledVersion() (string, error) {
	version, err := plist.ReadPlistString(internal.ChromeInfoPlistPath,
		"CFBundleShortVersionString")
	if err != nil {
		return "", fmt.Errorf("Could not get Chrome version because %s", err.Error())
	}
	return version, nil
}

// MajorVersion takes the InstalledVersion string and returns the major version as an int.
func MajorVersion() (int, error) {
	version, err := InstalledVersion()
	if err != nil {
		return 0, fmt.Errorf("failed to parse chrome major version: %s", err.Error())
	}

	// version is usually of the format 99.0.4844.83
	spl := strings.Split(version, ".")

	majorVersion, err := strconv.Atoi(spl[0])
	if err != nil {
		return 0, fmt.Errorf("failed to parse chrome major version: %s", err.Error())
	}
	return majorVersion, nil
}
