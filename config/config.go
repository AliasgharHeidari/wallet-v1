package config

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Server   `yaml:"server"`
	Database `yaml:"database"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	DSN DSNConfig `yaml:"dsn"`
}

type DSNConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
	TimeZone string `yaml:"timezone"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var Cfg Config

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		return nil, err
	}

	return &Cfg, nil

}

func (d DSNConfig) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		d.Host,
		d.User,
		d.Password,
		d.DBName,
		d.Port,
		d.SSLMode,
		d.TimeZone,
	)

}
