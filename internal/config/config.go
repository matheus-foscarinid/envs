package config

type Config struct {
	Version int    `json:"version"`
	Active  string `json:"active"`
}

func Load(path string) (Config, error) {
	return Config{}, nil
}

func Write(path string, data Config) error {
	return nil
}
