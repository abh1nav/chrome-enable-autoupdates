package checks

import (
	"fmt"
	"os"

	"github.com/abh1nav/chrome-enable-autoupdates/internal"
	"github.com/abh1nav/chrome-enable-autoupdates/internal/plist"
)

// Makes sure the `Google Chrome.app` exists at the expected location.
func EnsureChromeInstalled() error {
	if _, err := os.Stat(internal.ChromePath); os.IsNotExist(err) {
		return fmt.Errorf("Chrome is not installed on this computer")
	}
	return nil
}

// ChromeVersion reads the Info.plist and returns the value for CFBundleShortVersionString
// i.e. the installed version of Chrome.
func ChromeVersion() (string, error) {
	version, err := plist.ReadPlistString(internal.ChromeInfoPlistPath,
		"CFBundleShortVersionString")
	if err != nil {
		return "", fmt.Errorf("Could not get Chrome version because %s", err.Error())
	}
	return version, nil
}

// ChromeUpdateURL reads the Info.plist and returns the value for KSUpdateURL
func ChromeUpdateURL() (string, error) {
	updateURL, err := plist.ReadPlistString(internal.ChromeInfoPlistPath,
		"KSUpdateURL")
	if err != nil {
		return "", fmt.Errorf("Could not get Chrome update URL because %s", err.Error())
	}
	return updateURL, nil
}

// ChromeProductID reads the Info.plist and returns the value for KSProductID
func ChromeProductID() (string, error) {
	productID, err := plist.ReadPlistString(internal.ChromeInfoPlistPath,
		"KSProductID")
	if err != nil {
		return "", fmt.Errorf("Could not get Chrome Product ID because %s", err.Error())
	}
	return productID, nil
}
