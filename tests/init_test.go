package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitCreatesAllFiles(t *testing.T) {
	dir := t.TempDir()
	output, err := runEnvs(t, dir, "init")
	if err != nil {
		t.Fatalf("init failed: %v\noutput: %s", err, output)
	}

	for _, name := range []string{".envs.json", ".env.sample", ".env"} {
		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("expected %s to be created", name)
		}
	}
}

func TestInitConfigHasCorrectContent(t *testing.T) {
	dir := t.TempDir()
	if _, err := runEnvs(t, dir, "init"); err != nil {
		t.Fatalf("init failed: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, ".envs.json"))
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}

	var cfg struct {
		Version int    `json:"version"`
		Active  string `json:"active"`
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		t.Fatalf("invalid JSON in config: %v", err)
	}

	if cfg.Version != 1 {
		t.Errorf("expected version 1, got %d", cfg.Version)
	}
	if cfg.Active != "default" {
		t.Errorf("expected active \"default\", got %q", cfg.Active)
	}
}

func TestInitSkipsExistingFiles(t *testing.T) {
	dir := t.TempDir()

	content := []byte("existing content")
	for _, name := range []string{".envs.json", ".env.sample", ".env"} {
		os.WriteFile(filepath.Join(dir, name), content, 0644)
	}

	output, err := runEnvs(t, dir, "init")
	if err != nil {
		t.Fatalf("init failed: %v\noutput: %s", err, output)
	}

	for _, name := range []string{".envs.json", ".env.sample", ".env"} {
		data, _ := os.ReadFile(filepath.Join(dir, name))
		if string(data) != "existing content" {
			t.Errorf("expected %s to keep existing content, got %q", name, data)
		}
	}

	if !strings.Contains(output, "already exists") {
		t.Errorf("expected output to mention skipping existing files, got:\n%s", output)
	}
}
