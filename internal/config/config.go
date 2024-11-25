package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"development"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

// MustLoad loads and processes environment variables.
// Must means panic in case of non-standard execution of the function.
func MustLoad() *Config {
	// determine the source of environment variables
	config := os.Getenv("CONFIG_PATH")
	if config == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	// check if file exists
	if _, err := os.Stat(config); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", config)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(config, &cfg); err != nil {
		log.Fatalf("can't read config: %s", err)
	}

	return &cfg
}
