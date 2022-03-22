package main

import (
	"fmt"
	"os"

	"github.com/abh1nav/chrome-enable-autoupdates/internal/checks"
	"github.com/abh1nav/chrome-enable-autoupdates/internal/chrome"
	"github.com/abh1nav/chrome-enable-autoupdates/internal/keystone"
)

func throw(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func main() {
	// Make sure the process is being run as root
	err := checks.EnsureRoot()
	if err != nil {
		throw(err)
	}
	fmt.Println("Command running as root")

	// Ensure Chrome is installed
	err = chrome.EnsureInstalled()
	if err != nil {
		throw(err)
	}
	fmt.Println("Chrome is installed")

	chromeVersion, err := chrome.InstalledVersion()
	if err != nil {
		throw(err)
	}
	fmt.Printf("Chrome version is %s\n", chromeVersion)

	chromeUpdateURL, err := chrome.UpdateURL()
	if err != nil {
		throw(err)
	}
	fmt.Printf("Chrome update URL is %s\n", chromeUpdateURL)

	chromeProductID, err := chrome.ProductID()
	if err != nil {
		throw(err)
	}
	fmt.Printf("Chrome product ID is %s\n", chromeProductID)

	keystoneRegistrationFramework, err := keystone.RegistrationFrameworkPath()
	if err != nil {
		throw(err)
	}
	fmt.Printf("Keystone Registration Framework is %s\n", keystoneRegistrationFramework)
}
