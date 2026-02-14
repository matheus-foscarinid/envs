package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestListNoEnvFiles(t *testing.T) {
	dir := t.TempDir()
	output, err := runEnvs(t, dir, "list")
	if err != nil {
		t.Fatalf("list failed: %v\noutput: %s", err, output)
	}

	if !strings.Contains(output, "No envs found") {
		t.Errorf("expected 'No envs found', got:\n%s", output)
	}
}

func TestListShowsEnvFiles(t *testing.T) {
	dir := t.TempDir()

	os.WriteFile(filepath.Join(dir, ".env.dev"), []byte{}, 0644)
	os.WriteFile(filepath.Join(dir, ".env.prod"), []byte{}, 0644)

	output, err := runEnvs(t, dir, "list")
	if err != nil {
		t.Fatalf("list failed: %v\noutput: %s", err, output)
	}

	if !strings.Contains(output, ".env.dev") {
		t.Errorf("expected .env.dev in output, got:\n%s", output)
	}
	if !strings.Contains(output, ".env.prod") {
		t.Errorf("expected .env.prod in output, got:\n%s", output)
	}
}

func TestListExcludesSampleFile(t *testing.T) {
	dir := t.TempDir()

	os.WriteFile(filepath.Join(dir, ".env.sample"), []byte{}, 0644)
	os.WriteFile(filepath.Join(dir, ".env.dev"), []byte{}, 0644)

	output, err := runEnvs(t, dir, "list")
	if err != nil {
		t.Fatalf("list failed: %v\noutput: %s", err, output)
	}

	if !strings.Contains(output, ".env.dev") {
		t.Errorf("expected .env.dev in output, got:\n%s", output)
	}

	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == ".env.sample" || trimmed == "* .env.sample" {
			t.Errorf(".env.sample should be excluded from listing, but found: %q", line)
		}
	}
}
