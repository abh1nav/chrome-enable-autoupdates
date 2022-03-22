package main

import (
	"fmt"
	"os"

	"github.com/abh1nav/chrome-enable-autoupdates/internal/chrome"
	"github.com/abh1nav/chrome-enable-autoupdates/internal/keystone"
	"github.com/abh1nav/chrome-enable-autoupdates/internal/system"
)

var version string //nolint:unused

func throw(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func main() {
	// Print version
	fmt.Printf("chrome-enable-autoupdates v%s\n", version)
	fmt.Println("---")

	// Make sure the process is being run as root
	err := system.EnsureRoot()
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
	fmt.Printf("Found Chrome version %s\n", chromeVersion)

	fmt.Println("Reinstalling Keystone")
	err = keystone.Install()
	if err != nil {
		throw(err)
	}
	fmt.Println("")

	fmt.Println("Register Chrome with Keystone")
	err = chrome.RegisterWithKeystone()
	if err != nil {
		throw(err)
	}
}
