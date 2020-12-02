package app

import "github.com/ednailson/serasa-challenge/database"

type Config struct {
	Key          string          `json:"key"`
	Nonce        string          `json:"nonce"`
	MainframeUrl string          `json:"mainframe_url"`
	Port         int             `json:"port"`
	Database     database.Config `json:"database"`
}
