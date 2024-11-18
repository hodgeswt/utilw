package main

import (
	"log"
	"os"

	"github.com/hodgeswt/utilw/internal/catw"
)

func main() {
	if err := catw.Run(os.Args); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
