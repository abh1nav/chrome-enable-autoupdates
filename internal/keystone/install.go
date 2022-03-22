package keystone

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/abh1nav/chrome-enable-autoupdates/internal/chrome"
	"github.com/abh1nav/chrome-enable-autoupdates/internal/system"
)

// Install forces keystone to reinstall itself.
// If a non-python installer is not found, this will throw an error
// because we're only trying to support MacOS 12.3 and above without python.
func Install() error {
	majorVersion, err := chrome.MajorVersion()
	if err != nil {
		return err
	}

	registrationPath, err := RegistrationFrameworkPath()
	if err != nil {
		return err
	}

	installScript := filepath.Join(registrationPath, "Resources", "ksinstall")
	if majorVersion >= 80 {
		installScript = filepath.Join(registrationPath, "Helpers", "ksinstall")
	}
	if !system.Exists(installScript) {
		return fmt.Errorf("Keystone install script %s does not exist", installScript)
	}

	installPayload := filepath.Join(registrationPath, "Resources", "Keystone.tbz")
	if !system.Exists(installPayload) {
		return fmt.Errorf("Keystone install payload %s does not exist", installPayload)
	}

	installOutput, err := execInstall(installScript, installPayload)
	if err != nil {
		return err
	}

	fmt.Println(">> Keystone Install Output")
	fmt.Println(installOutput)
	fmt.Println(">> End Keystone Install Output")
	return nil
}

// execInstall runs the command to install keystone
func execInstall(installScript, installPayload string) (string, error) {
	stdout, err := exec.Command(installScript, "--install", installPayload, "--force").Output()
	if err != nil {
		return "", err
	}
	return string(stdout), nil
}
