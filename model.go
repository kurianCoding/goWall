package main

import "fmt"

func ChangeWallPaper(path string, conn *sql.Conn) {
	if valid := validate(path); valid {
		conn.Write(fmt.Sprintf("UPDATE DATA SET VALUE = %s", path))
	}
	return
}

func ResetDock() {
	// killall Dock
}

func validate(string path) {
	return true
}
