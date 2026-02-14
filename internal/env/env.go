package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Load parses a .env file into a key-value map
func Load(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	envMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		envMap[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}

	return envMap, scanner.Err()
}

// Write serializes a key-value map into a .env file, preserving comments and blank lines
func Write(path string, data map[string]string) error {
	// read existing file to preserve comments and order
	existing, err := os.Open(path)
	if err != nil {
		return writeFromScratch(path, data)
	}
	defer existing.Close()

	written := make(map[string]bool)
	var lines []string
	scanner := bufio.NewScanner(existing)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// preserve comments and blank lines
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			lines = append(lines, line)
			continue
		}

		key, _, found := strings.Cut(trimmed, "=")
		if !found {
			lines = append(lines, line)
			continue
		}

		key = strings.TrimSpace(key)
		if val, ok := data[key]; ok {
			lines = append(lines, fmt.Sprintf("%s=%s", key, val))
			written[key] = true
		} else {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// append new keys not in original file
	for key, val := range data {
		if !written[key] {
			lines = append(lines, fmt.Sprintf("%s=%s", key, val))
		}
	}

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}

func writeFromScratch(path string, data map[string]string) error {
	var lines []string
	for key, val := range data {
		lines = append(lines, fmt.Sprintf("%s=%s", key, val))
	}
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}

func ListEnvs() ([]string, error) {
	return nil, nil
}

func Copy(source string, destination string) error {
	return nil
}
