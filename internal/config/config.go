package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port       string `yaml:"port" envconfig:"SERVER_PORT"`
		Host       string `yaml:"host" envconfig:"SERVER_HOST"`
		SessionKey string `yaml:"session_key" envconfig:"SESSION_KEY"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user" envconfig:"DB_USERNAME"`
		Password string `yaml:"pass" envconfig:"DB_PASSWORD"`
		File     string `yaml:"file" envconfig:"DB_FILE"`
	} `yaml:"database"`
	Frontend struct {
		Admin    string `yaml:"admin" envconfig:"USERNAME"`
		Password string `yaml:"pass" envconfig:"PASSWORD"`
		SiteName string `yaml:"site_name" envconfig:"FILE"`
	} `yaml:"frontend"`
}

var Cfg Config

func init() {
	readFile(&Cfg)
	readEnv(&Cfg)
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		panic(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		panic(err)
	}
}
