package main

import "os"

func main() {
	config, err := getNewConfig()
	if err != nil {
		os.Exit(1)
	}
	// function that will change the wallpaper
	ChangeWallPaper(config.NewWallpaper, config.Conn)
	ResetDock()
	// function that will loop the application with a timer
	return
}
