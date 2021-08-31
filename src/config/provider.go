package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type JWT struct {
	SecretKey string `yaml:"secret_key"`
}

type Database struct {
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Net string `yaml:"net"`
	Host string `yaml:"host"`
	Dbname string `yaml:"dbname"`
	AllowNativePasswords bool `yaml:"allow_native_passwords"`
}

type Config struct {
	Database `yaml:"database"`
	JWT `yaml:"jwt"`
}

func GetConfig() (*Config, error)  {
	f, err := os.Open("config.yaml")

	if err != nil {
		// todo: handle
		fmt.Println("Error in os.Open: " + err.Error())

		return &Config{}, err
	}

	defer f.Close()

	var cfg Config

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		//todo: handle
		fmt.Println("Error in decoder.Decode: " + err.Error())

		return &Config{}, err
	}

	return &cfg, nil
}
