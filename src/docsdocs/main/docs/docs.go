package main

import (
	"docsdocs/log"
)

func main() {
	log.LogTo("json", "stdout", "debug")
	logTest := log.NewDocsLogger()
	logTest.Debug("DocsDocs Crypt Debug Log")
}
