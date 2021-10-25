package main

import (
	"fmt"
	"log"

	"github.com/KeisukeYamashita/go-macos-keylogger/pkg/keyboard"
	"github.com/KeisukeYamashita/go-macos-keylogger/pkg/keylogger"
)

func main() {
	kl, err := keylogger.New()
	if err != nil {
		log.Fatal(err)
	}

	f := func(key keyboard.Key, state keyboard.State) {
		fmt.Println(key)
	}

	kl.Listen(f)
}
