package main

import (
	"fmt"
	"os"

	"github.com/abh1nav/chrome-enable-autoupdates/internal/checks"
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
	err = checks.EnsureChromeInstalled()
	if err != nil {
		throw(err)
	}
	fmt.Println("Chrome is installed")

	chromeVersion, err := checks.ChromeVersion()
	if err != nil {
		throw((err))
	}
	fmt.Printf("Chrome version is %s\n", chromeVersion)

	chromeUpdateURL, err := checks.ChromeUpdateURL()
	if err != nil {
		throw((err))
	}
	fmt.Printf("Chrome update URL is %s\n", chromeUpdateURL)

	chromeProductID, err := checks.ChromeProductID()
	if err != nil {
		throw((err))
	}
	fmt.Printf("Chrome product ID is %s\n", chromeProductID)
}
