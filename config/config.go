package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	ServerURI      string `json:"serverURI"`
	LocalPort      int    `json:"localPort"`
	AudioPath      string `json:"audioPath"`
	MediaFilesRoot string `json:"mediaFilesRoot"`
	ValidTokenStr  string `json:"validTokenStr"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
