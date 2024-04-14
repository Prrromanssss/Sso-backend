package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env" env-default:"local"`
	StorageUrl string     `yaml:"storage_url" env-default:"postgres://postgres:postgres@localhost:5432/sso?sslmode=disable"`
	TokenTTL   string     `yaml:"token_ttl" env-required:"true"`
	GRPC       GRPCConfig `yaml:"grpc" env-required:"true"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-default:"8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	// --config="path/to/config.yaml"
	flag.StringVar(&res, "config", "", "path to config path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
