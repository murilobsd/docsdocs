package main

import (
	"docsdocs/client/views/gui"
	"docsdocs/log"
)

func main() {
	log.Settings("json", "stdout", "debug")
	logTest := log.NewDocsLogger()
	logTest.Debug("DocsDocs Crypt Debug Log")
	win := gui.NewGuiView()
	win.Run()
}
