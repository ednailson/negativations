package app

import (
	"encoding/json"
	"github.com/ednailson/serasa-challenge/database"
	. "github.com/ednailson/serasa-challenge/helper_tests"
	"os"
)

type Config struct {
	Key          string          `json:"key"`
	Nonce        string          `json:"nonce"`
	MainframeUrl string          `json:"mainframe_url"`
	Port         int             `json:"port"`
	Database     database.Config `json:"database"`
}

func NewConfigFile(filename string) error {
	err := generateConfigFile(filename, configSample())
	if err != nil {
		return err
	}
	return nil
}

func generateConfigFile(filename string, config Config) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func configSample() Config {
	return Config{
		Key:          Key,
		Nonce:        Nonce,
		MainframeUrl: "http://mainframe.service.com.br:3000/negativations",
		Port:         5000,
		Database: database.Config{
			Collection: "negativations",
			Host:       DBHostTest,
			Port:       DBPortTest,
			User:       DBUserTest,
			Password:   DBPassTest,
			Database:   DBNameTest,
		},
	}
}
