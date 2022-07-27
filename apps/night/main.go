package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ericluj/egit/config"
	"github.com/ericluj/egit/git"
	log "github.com/ericluj/elog"
)

func main() {
	st := time.Now().Unix()

	commits, err := git.GitLog()
	if err != nil {
		log.Fatalf("GitLog error: %v", err)
	}

	for _, v := range commits {
		replaceDt := replaceHour(v.AuthorDate)
		if v.AuthorDate == replaceDt {
			continue
		}

		err := git.EditCommit(v.Hash, replaceDt, config.Author, config.Email)
		if err != nil {
			log.Fatalf("EditCommit error: %v", err)
		}
	}

	ed := time.Now().Unix()
	time := ed - st
	log.Infof("egit night success! time: %ds", time)
}

func replaceHour(dt string) string {
	dtArr := strings.Split(dt, " ")
	weekDay := dtArr[0]
	month := dtArr[1]
	day := dtArr[2]
	ti := dtArr[3]
	year := dtArr[4]
	tiZone := dtArr[5]

	tiArr := strings.Split(ti, ":")
	hour := tiArr[0]
	minute := tiArr[1]
	second := tiArr[2]

	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		return dt
	}

	if hourInt < 19 {
		hourInt += 5
	}

	replaceTi := fmt.Sprintf("%d:%s:%s", hourInt, minute, second)

	replaceDt := fmt.Sprintf("%s %s %s %s %s %s", weekDay, month, day, replaceTi, year, tiZone)

	return replaceDt
}
