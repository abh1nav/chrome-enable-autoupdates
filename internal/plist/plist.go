package plist

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ReadPlistString looks for a key within a Plist file and returns its string value.
// This uses a brute force / brittle way of parsing Plists.
func ReadPlistString(filePath, key string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// We're looking for a line that looks like <key>{key}</key>
	lookingFor := fmt.Sprintf("<key>%s</key>", key)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, lookingFor) {
			// We've found the key we're looking for.
			// Expecting the next line to be the value we need to return.
			if scanner.Scan() {
				return parseValue(scanner.Text())
			}
			if err = scanner.Err(); err != nil {
				return "", err
			}
		}
	}
	if err = scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("Key %s not found in %s", key, filePath)
}

// parseValue takes a line of the format `<string>Value</string>` and returns the `Value`.
// `line` may have leading and/or trailing spaces.
func parseValue(line string) (string, error) {
	if !strings.Contains(line, "<string>") {
		return "", errors.New("Value is not a string")
	}

	split1 := strings.Split(line, "<string>")
	if len(split1) != 2 {
		return "", fmt.Errorf("Malformed value from plist: %s", line)
	}

	split2 := strings.Split(split1[1], "</string>")
	if len(split2) != 2 {
		return "", fmt.Errorf("Malformed value from plist: %s", line)
	}

	return split2[0], nil
}
