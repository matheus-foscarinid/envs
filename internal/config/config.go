package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/matheus-foscarinid/envs/internal/constants"
)

type Config struct {
	Version int    `json:"version"`
	Active  string `json:"active"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func Write(path string, data Config) error {
	file, err := os.Create(constants.ConfigFile)
	if err != nil {
		fmt.Println("Error creating config file:", err)
		return err
	}
	defer file.Close()

	json.NewEncoder(file).Encode(data)
	return nil
}
