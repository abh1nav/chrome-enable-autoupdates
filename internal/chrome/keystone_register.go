package chrome

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/abh1nav/chrome-enable-autoupdates/internal"
)

func RegisterWithKeystone() error {
	ksadmin := filepath.Join("/Library", "Google", "GoogleSoftwareUpdate",
		"GoogleSoftwareUpdate.bundle", "Contents", "MacOS", "ksadmin")

	productID, err := ProductID()
	if err != nil {
		return err
	}

	version, err := InstalledVersion()
	if err != nil {
		return err
	}

	updateURL, err := UpdateURL()
	if err != nil {
		return err
	}

	stdout, err := exec.Command(ksadmin, "--register",
		"--productid", productID,
		"--version", version,
		"--xcpath", internal.ChromePath,
		"--url", updateURL,
		"--tag-path", internal.ChromeInfoPlistPath,
		"--tag-key", internal.ChromeTagKey,
		"--brand-path", internal.ChromeBrandPath,
		"--brand-key", internal.ChromeBrandKey,
		"--version-path", internal.ChromeInfoPlistPath,
		"--version-key", internal.ChromeVersionKey).
		Output()

	if err != nil {
		return err
	}

	fmt.Println(">> Chrome Keystone Register Output")
	fmt.Println(string(stdout))
	fmt.Println(">> End Chrome Keystone Register Output")

	return nil
}
