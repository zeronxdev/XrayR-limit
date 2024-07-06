package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/zeronxdev/XrayR-limit/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
