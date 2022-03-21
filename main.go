package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"howett.net/plist"
)

const CHROME_PATH = "/Applications/Google Chrome.app"

var CHROME_INFO_PLIST_PATH = path.Join(CHROME_PATH, "Contents/Info.plist")

type ChromeInfoPList struct {
	ShortVersion string `plist:"CFBundleShortVersionString"`
	UpdateURL    string `plist:"KSUpdateURL"`
	ProductID    string `plist:"KSProductID"`
}

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

func parseInfoPList() (ChromeInfoPList, error) {
	fmt.Println("Reading plist", CHROME_INFO_PLIST_PATH)
	var p ChromeInfoPList
	f, err := os.Open(CHROME_INFO_PLIST_PATH)
	if err != nil {
		return p, err
	}
	defer f.Close()
	var i interface{}
	err = plist.NewDecoder(f).Decode(i)
	if err != nil {
		fmt.Printf("Failed to read plist: %s\n", err.Error())
		return p, err
	}
	fmt.Printf("plist is %+v\n", i)
	return p, err
}

func getChromeVersion() string {
	pList, err := parseInfoPList()
	if err != nil {
		fmt.Printf("Error: failed to open Chrome info plist %s because: %s", CHROME_INFO_PLIST_PATH, err.Error())
		os.Exit(1)
	}
	return pList.ShortVersion
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
