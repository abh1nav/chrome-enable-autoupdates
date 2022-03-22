package keystone

import (
	"path/filepath"

	"github.com/abh1nav/chrome-enable-autoupdates/internal"
	"github.com/abh1nav/chrome-enable-autoupdates/internal/chrome"
)

// RegistrationFrameworkPath returns the path to the Keystone Registration Framework
func RegistrationFrameworkPath() (string, error) {
	fullVersion, err := chrome.InstalledVersion()
	if err != nil {
		return "", err
	}

	majorVersion, err := chrome.MajorVersion()
	if err != nil {
		return "", err
	}

	var path string
	if majorVersion >= 76 {
		path = filepath.Join(internal.ChromePath, "Contents", "Frameworks", "Google Chrome Framework.framework", "Frameworks", "KeystoneRegistration.framework", "Versions", "Current")
	} else if majorVersion >= 75 {
		path = filepath.Join(internal.ChromePath, "Contents", "Frameworks", "Google Chrome Framework.framework", "Versions", fullVersion, "Frameworks/KeystoneRegistration.framework")
	} else {
		path = filepath.Join(internal.ChromePath, "Contents", "Versions", fullVersion, "Google Chrome Framework.framework", "Frameworks", "KeystonRegistration.framework")
	}

	return path, nil
}
