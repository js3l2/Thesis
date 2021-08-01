package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Started...")
	for i := 0; i < 3; i++ {
		log.Warn("counting...", i)
	}
	log.Warn("hello")
}
