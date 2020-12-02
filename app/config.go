package app

import "github.com/ednailson/serasa-challenge/database"

type Config struct {
	MainframeUrl string          `json:"mainframe_url"`
	Port         int             `json:"port"`
	Database     database.Config `json:"database"`
}
