package manager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matheus-foscarinid/envs/internal/config"
	"github.com/matheus-foscarinid/envs/internal/constants"
)

type Manager struct {
	cfg *config.Config
}

// Init sets up envs in the current directory, creating config, sample, and env files
func Init() error {
	fmt.Println("⛰️ Initializing envs...")
	fmt.Println()

	createIfMissing(constants.ConfigFile, "config", func() error {
		defaultCfg := config.Config{Version: 1, Active: "default"}
		return config.Write(constants.ConfigFile, defaultCfg)
	})
	createIfMissing(constants.SampleFile, "sample", func() error {
		_, err := os.Create(constants.SampleFile)
		return err
	})
	createIfMissing(constants.DotEnvFile, "env", func() error {
		_, err := os.Create(constants.DotEnvFile)
		return err
	})

	fmt.Println()
	fmt.Println("⛰️ Envs initialized successfully!")
	return nil
}

func createIfMissing(path string, label string, create func() error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Creating default %s file...\n", label)
		if err := create(); err != nil {
			fmt.Printf("Error creating %s file: %v\n", label, err)
		}
	} else {
		fmt.Printf("✅ %s file already exists, skipping...\n", label)
	}
}

// NewEnv creates a new environment file, optionally copying from the sample file
func NewEnv(name string, empty bool) error {
	fmt.Println("⛰️ Creating new environment:", name)

	envFile := fmt.Sprintf(".env.%s", name)

	if _, err := os.Stat(envFile); !os.IsNotExist(err) {
		fmt.Printf("environment file %s already exists\nskipping...\n", name)
		return nil
	}

	if _, err := os.Create(envFile); err != nil {
		return fmt.Errorf("error creating environment file: %w", err)
	}

	if !empty {
		if _, err := os.Stat(constants.SampleFile); os.IsNotExist(err) {
			fmt.Println("sample file not found, creating empty environment file...")
		} else {
			sampleContent, err := os.ReadFile(constants.SampleFile)
			if err != nil {
				return fmt.Errorf("error reading sample file: %w", err)
			}
			if err := os.WriteFile(envFile, sampleContent, 0644); err != nil {
				return fmt.Errorf("error writing environment file: %w", err)
			}
		}
	}

	fmt.Printf("✅ .env.%s created successfully\n", name)
	return nil
}

func (m *Manager) Use(name string) error {
	return nil
}
func (m *Manager) Current() error {
	return nil
}
func (m *Manager) Set(key,value string) error {
	return nil
}
func (m *Manager) Get(key string) error {
	return nil
}
func (m *Manager) Validate(name string) error {
	return nil
}

func ListEnvs() ([]string, error) {
	pattern := ".env.*"

	files, err := filepath.Glob(pattern)
	filteredFiles := []string{}
	for _, file := range files {
		if file == constants.SampleFile {
			continue
		}
		filteredFiles = append(filteredFiles, file)
	}

	cfg, err := config.Load(constants.ConfigFile)
	if err != nil {
		return nil, err
	}
	currentEnv := cfg.Active

	if len(filteredFiles) == 0 {
		fmt.Println("No envs found")
		return nil, nil
	}

	for _, file := range filteredFiles {
		if file == currentEnv {
			fmt.Println("*", file)
		} else {
			fmt.Println(file)
		}
	}

	return nil, nil
}
