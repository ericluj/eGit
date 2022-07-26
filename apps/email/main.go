package main

import (
	"time"

	"github.com/ericluj/egit/config"
	"github.com/ericluj/egit/git"
	log "github.com/ericluj/elog"
)

func main() {
	st := time.Now().Unix()

	commits, err := git.GitLog()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, v := range commits {
		err := git.EditCommit(v.Hash, v.AuthorDate, config.Author, config.Email)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}

	ed := time.Now().Unix()
	time := ed - st
	log.Infof("egit email success! time: %ds", time)
}
