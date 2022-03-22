package chrome

import (
	"fmt"

	"github.com/abh1nav/chrome-enable-autoupdates/internal"
	"github.com/abh1nav/chrome-enable-autoupdates/internal/plist"
	"github.com/abh1nav/chrome-enable-autoupdates/internal/system"
)

// Makes sure the `Google Chrome.app` exists at the expected location.
func EnsureInstalled() error {
	if !system.Exists(internal.ChromePath) {
		return fmt.Errorf("Chrome not found at %s", internal.ChromePath)
	}
	return nil
}

// UpdateURL reads the Info.plist and returns the value for KSUpdateURL
func UpdateURL() (string, error) {
	updateURL, err := plist.ReadPlistString(internal.ChromeInfoPlistPath,
		"KSUpdateURL")
	if err != nil {
		return "", fmt.Errorf("Could not get Chrome update URL because %s", err.Error())
	}
	return updateURL, nil
}

// ProductID reads the Info.plist and returns the value for KSProductID
func ProductID() (string, error) {
	productID, err := plist.ReadPlistString(internal.ChromeInfoPlistPath,
		"KSProductID")
	if err != nil {
		return "", fmt.Errorf("Could not get Chrome Product ID because %s", err.Error())
	}
	return productID, nil
}
