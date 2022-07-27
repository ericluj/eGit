package main

import (
	"time"

	"github.com/ericluj/egit/config"
	"github.com/ericluj/egit/git"
	log "github.com/ericluj/elog"
)

func main() {
	st := time.Now().Unix()

	err := git.EditAllCommit(config.Author, config.Email)
	if err != nil {
		log.Fatalf("EditAllCommit error: %v", err)
	}

	ed := time.Now().Unix()
	time := ed - st
	log.Infof("egit email success! time: %ds", time)
}
