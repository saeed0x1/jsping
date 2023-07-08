package main

import (
	"log"

	"github.com/saeed0x1/jsping/runner/jsping"
)

func main() {
	opts := jsping.ParseOptions()
	runner := jsping.New(opts)
	err := runner.Run()
	if err != nil {
		log.Fatalf("Error running jsping: %s", err)
	}
}
