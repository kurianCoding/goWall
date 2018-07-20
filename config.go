package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/slicebit/qb"
	"os"
)

type Config struct {
	NewWallpaper string
	Conn         *mysql.Conn
}

// get config
func getNewConfig() (*Config, error) {
	db, err := qb.New("sqlite3", "~/Library/Application Support/Dock/desktoppicture.db")
	if err != nil {
		os.Exit(err)
	}
	return nil, &Config{}
}

// get config from envs
func getEnvs() (error, map[string]string) {}
