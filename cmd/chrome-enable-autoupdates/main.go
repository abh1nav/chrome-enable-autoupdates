package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/abh1nav/chrome-enable-autoupdates/internal/plist"
)

const CHROME_PATH = "/Applications/Google Chrome.app"

var CHROME_INFO_PLIST_PATH = path.Join(CHROME_PATH, "Contents/Info.plist")

func ensureRoot() {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
		fmt.Printf("Error: Failed to determine owner of process: %s\n", err.Error())
		os.Exit(1)
	}

	owner := string(stdout)
	owner = strings.TrimSuffix(owner, "\n")
	if owner != "root" {
		fmt.Printf("This script must be run as root, current user: '%s'\n", owner)
		os.Exit(1)
	}
}

func ensureChromeInstalled() {
	if _, err := os.Stat(CHROME_PATH); os.IsNotExist(err) {
		fmt.Printf("Error: Chrome is not installed on this computer\n")
		os.Exit(1)
	}
}

func getChromeVersion() string {
	version, err := plist.ReadPlistString(CHROME_INFO_PLIST_PATH, "CFBundleShortVersionString")
	if err != nil {
		fmt.Printf("Error: failed to open Chrome info plist %s because: %s", CHROME_INFO_PLIST_PATH, err.Error())
		os.Exit(1)
	}
	return version
}

func main() {
	// Make sure the process is being run as root
	ensureRoot()
	fmt.Println("Command running as root")

	// Ensure Chrome is installed
	ensureChromeInstalled()
	chromeVersion := getChromeVersion()
	fmt.Printf("Chrome version %s is installed at %s\n", chromeVersion, CHROME_PATH)
}
