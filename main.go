package main

import (
	"log"

	"github.com/gen2brain/beeep"
)

func main() {
	err := beeep.Notify("Title", "Message body", "assets/information.png")
	if err != nil {
		log.Panic(err)
	}
}
