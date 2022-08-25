package main

import (
	"time"

	"github.com/ericluj/egit/git"
	log "github.com/ericluj/elog"
)

func main() {
	st := time.Now().Unix()

	hash := ""
	comment := ""

	err := git.EditComment(hash, comment)
	if err != nil {
		log.Fatalf("EditComment error: %v", err)
	}

	ed := time.Now().Unix()
	time := ed - st
	log.Infof("egit comment success! time: %ds", time)
}
