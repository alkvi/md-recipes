package main

import (
    "gopkg.in/yaml.v3"
    "os"
)

type AppConfig struct {
    FolderPath string `json:"folderPath" yaml:"folder_path"`
    LogLevel   string `json:"logLevel" yaml:"log_level"`
}

func LoadConfig(path string) (*AppConfig, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var config AppConfig
    decoder := yaml.NewDecoder(file)
    err = decoder.Decode(&config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}
