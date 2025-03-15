package main

import (
	"log"

	"Github.com/Yobubble/email-virus-scanner-server/config"
)

func main() {
	cfg := config.InitConfig()
	log.Printf("%v", cfg)
}
