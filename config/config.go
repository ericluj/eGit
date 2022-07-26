package config

import (
	log "github.com/ericluj/elog"
)

const (
	Author      = ""
	Email       = ""
	ProjectPath = ""
)

func init() {
	if Author == "" || Email == "" || ProjectPath == "" {
		log.Fatalf("config invalid")
	}
}
