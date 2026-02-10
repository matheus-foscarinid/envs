package manager

import (
	"fmt"
	"path/filepath"

	"github.com/matheus-foscarinid/envs/internal/config"
	"github.com/matheus-foscarinid/envs/internal/constants"
)

type Manager struct {
	cfg *config.Config
}

func currentEnv() (string, error) {
	// get the active env from the config file
	cfg, err := config.Load(constants.ConfigFile)
	if err != nil {
		return "", err
	}
	return cfg.Active, nil
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

	currentEnv, err := currentEnv()
	if err != nil {
		return nil, err
	}

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
