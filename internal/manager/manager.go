package manager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matheus-foscarinid/envs/internal/config"
	"github.com/matheus-foscarinid/envs/internal/constants"
	"github.com/matheus-foscarinid/envs/internal/env"
)

type Manager struct {
	cfg *config.Config
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

func copyFile(source, destination string) error {
	sourceFile, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	return os.WriteFile(destination, sourceFile, 0644)
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
			if err := copyFile(constants.SampleFile, envFile); err != nil {
				return fmt.Errorf("error writing environment file: %w", err)
			}
		}
	}

	fmt.Printf("✅ .env.%s created successfully\n", name)
	return nil
}

func Use(name string) error {
	fmt.Println("⛰️ Switching active environment to:", name)

	cfg, err := config.Load(constants.ConfigFile)
	if err != nil {
		return err
	}
	cfg.Active = name

	err = copyFile(fmt.Sprintf(".env.%s", name), constants.DotEnvFile)
	if err != nil {
		return err
	}

	fmt.Println("✅ Active environment switched to:", name)
	return nil
}

func Current() error {
	cfg, err := config.Load(constants.ConfigFile)
	if err != nil {
		return err
	}
	currentEnv := cfg.Active
	fmt.Println("✅ Active environment:", currentEnv)
	return nil
}

// Set updates or adds a key-value pair in the active .env file
func Set(key, value string) error {
	envMap, err := env.Load(constants.DotEnvFile)
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	envMap[key] = value
	if err := env.Write(constants.DotEnvFile, envMap); err != nil {
		return fmt.Errorf("error writing .env file: %w", err)
	}

	fmt.Printf("✅ %s=%s\n", key, value)
	return nil
}

// Get prints the value of a key from the active .env file
func Get(key string) error {
	envMap, err := env.Load(constants.DotEnvFile)
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	val, ok := envMap[key]
	if !ok {
		fmt.Printf("key %q not found in .env\n", key)
		return nil
	}

	fmt.Println(val)
	return nil
}
// Validate checks that every key in .env.sample is present in the active .env
func Validate() error {
	sampleVars, err := env.Load(constants.SampleFile)
	if err != nil {
		return fmt.Errorf("could not load %s: %w", constants.SampleFile, err)
	}

	envVars, err := env.Load(constants.DotEnvFile)
	if err != nil {
		return fmt.Errorf("could not load %s: %w", constants.DotEnvFile, err)
	}

	var missing []string
	for key := range sampleVars {
		if _, ok := envVars[key]; !ok {
			missing = append(missing, key)
		}
	}

	if len(missing) == 0 {
		fmt.Println("✅ All sample keys are set in .env")
		return nil
	}

	fmt.Printf("⚠️  Missing %d key(s) in %s:\n", len(missing), constants.DotEnvFile)
	for _, key := range missing {
		fmt.Printf("  - %s\n", key)
	}
	return fmt.Errorf("validation failed: %d missing key(s)", len(missing))
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
		if file == fmt.Sprintf(".env.%s", currentEnv) {
			fmt.Println("*", file)
		} else {
			fmt.Println(" ", file)
		}
	}

	return nil, nil
}
