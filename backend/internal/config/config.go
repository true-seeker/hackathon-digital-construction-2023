package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Ujin       `yaml:"ujin"`
	Storage    `yaml:"storage"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	Port        string        `yaml:"port"`
}

type Ujin struct {
	ApiUrl     string `yaml:"api_url"`
	UserToken  string `yaml:"user_token"`
	AdminToken string `yaml:"admin_token"`
}

type Storage struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	Database string `yaml:"database"`
}

var Cfg Config

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/local.yaml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	isContainer := os.Getenv("IS_CONTAINER")
	if isContainer != "" {
		cfg.Storage.Address = "database"
		cfg.HTTPServer.Address = "webapi"
	}
	Cfg = cfg
	return &cfg
}
