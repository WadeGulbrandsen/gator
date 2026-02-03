package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DB_URL    string `json:"db_url"`
	User_Name string `json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}
	cfg := Config{}
	if err := json.Unmarshal(content, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c *Config) SetUser(user_name string) error {
	c.User_Name = user_name
	return write(c)
}

func write(cfg *Config) error {
	json_data, err := json.Marshal(*cfg)
	if err != nil {
		return err
	}
	file_path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	return os.WriteFile(file_path, json_data, 0666)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + "/" + configFileName, nil
}
