package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Estrutura para mapear o arquivo JSON
type Config struct {
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

// Função para carregar o arquivo de configuração
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("could not decode config JSON: %w", err)
	}

	return &cfg, nil
}
