package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const fileName = ".clc-term-config"

type Config struct {
	Username              string
	Password              string
	BearerToken           string
	BearerTokenExpiration time.Time
}

func (c Config) IsExpired() bool {
	return !time.Now().Before(c.BearerTokenExpiration)
}

func Read() (config string, err error) {
	homeDir, osErr := os.UserHomeDir()
	if osErr != nil {
		err = fmt.Errorf("failed to read user home directory")
		return
	}

	data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/%s", homeDir, fileName))
	if readErr != nil {
		err = fmt.Errorf("failed to read file %s/%s", homeDir, fileName)
		return
	}

	config = string(data)
	return
}

func Load() (config Config, err error) {
	data, err := Read()
	if err != nil {
		return
	}

	jErr := json.Unmarshal([]byte(data), &config)
	if jErr != nil {
		err = fmt.Errorf("failed to convert json to struct")
	}

	return
}

func Write(config Config) error {
	homeDir, osErr := os.UserHomeDir()
	if osErr != nil {
		return fmt.Errorf("failed to read user home directory")
	}

	filePath := fmt.Sprintf("%s/%s", homeDir, fileName)
	data, jErr := json.Marshal(config)
	if jErr != nil {
		return fmt.Errorf("failed to marshal config to json")
	}

	if err := ioutil.WriteFile(filePath, data, 0600); err != nil {
		return fmt.Errorf("failed to write config file")
	}

	return nil
}
