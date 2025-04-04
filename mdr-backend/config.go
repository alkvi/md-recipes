package main

import (
    "gopkg.in/yaml.v3"
    "os"
)

type AppConfig struct {
    FolderPath string `yaml:"folder_path"`
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
