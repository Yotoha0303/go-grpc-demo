package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	MySQL  MySQLConfig  `yaml:"mysql"`
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type MySQLConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

// DSN builds a MySQL DSN for GORM / go-sql-driver.
func (m MySQLConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.Database)
}

// LoadEnv loads optional .env into process environment.
// Missing .env is ignored so local runs can rely on shell env only.
func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		if _, statErr := os.Stat(".env"); os.IsNotExist(statErr) {
			return nil
		}
		return fmt.Errorf("load .env: %w", err)
	}
	return nil
}

func (Config) validate() error {

	return nil
}

// LoadConfig reads YAML config from path and fills password from env when empty.
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config %s: %w", path, err)
	}

	if cfg.MySQL.Password == "" {
		cfg.MySQL.Password = os.Getenv("MYSQL_PASSWORD")
	}
	if cfg.MySQL.User == "" {
		if u := os.Getenv("MYSQL_USER"); u != "" {
			cfg.MySQL.User = u
		}
	}
	if cfg.MySQL.Host == "" {
		if h := os.Getenv("MYSQL_HOST"); h != "" {
			cfg.MySQL.Host = h
		}
	}
	if cfg.MySQL.Database == "" {
		if d := os.Getenv("MYSQL_DATABASE"); d != "" {
			cfg.MySQL.Database = d
		}
	}

	return &cfg, nil
}
