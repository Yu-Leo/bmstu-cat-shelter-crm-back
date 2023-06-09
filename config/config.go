package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	pathToConfigFile = "./config/config.yaml" // Path to config file from base project dir
)

type Config struct {
	App     AppConfig     `yaml:"app"`
	Server  ServerConfig  `yaml:"server"`
	Logger  LoggerConfig  `yaml:"logger"`
	Storage StorageConfig `yaml:"storage"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type ServerConfig struct {
	Host string `env-required:"true" yaml:"host" env:"APP_HOST"`
	Port int    `env-required:"true" yaml:"port" env:"APP_PORT"`
	Mode string `yaml:"mode" env:"APP_MODE"` // "debug" or "release"
}

type LoggerConfig struct {
	Level string `yaml:"log_level" env:"LOG_LEVEL" env-default:"info"` // trace/debug/info/warning/error/fatal/panic
}

type StorageConfig struct {
	Path string `env-required:"true" yaml:"path" env:"DB_PATH"`
}

var instance *Config
var err error

var once sync.Once

func GetConfig() (*Config, error) {
	once.Do(func() {
		instance = &Config{}
		err = cleanenv.ReadConfig(pathToConfigFile, instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
		}
	})
	if instance == nil {
		return nil, err
	}
	return instance, nil
}
