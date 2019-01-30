package main

import (
	"docsdocs/log"
)

func main() {
	log.LogTo("text", "stdout", "debug")
	log.Info("DocsDocs Crypto Doc")
}
