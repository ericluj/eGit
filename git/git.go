package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/ericluj/egit/config"
)

type Commit struct {
	Hash       string
	Author     string
	AuthorDate string
	Commit     string
	CommitDate string
	Comment    string
}

func Cmd(name string, args ...string) (string, error) {
	c := exec.Command(name, args...)
	c.Dir = config.ProjectPath
	c.Env = []string{"FILTER_BRANCH_SQUELCH_WARNING=1"}
	out, err := c.CombinedOutput()
	return string(out), err
}

func getValue(str string) string {
	arr := strings.Split(str, ":")

	if len(arr) == 0 {
		return ""
	} else if len(arr) == 1 {
		return strings.TrimSpace(arr[0])
	} else {
		return strings.TrimSpace(arr[1])
	}
}

func getDate(str string) string {
	arr := strings.Split(str, "Date:")

	if len(arr) == 0 {
		return ""
	} else if len(arr) == 1 {
		return strings.TrimSpace(arr[0])
	} else {
		return strings.TrimSpace(arr[1])
	}
}

func GitLog() ([]*Commit, error) {
	commits := make([]*Commit, 0)
	out, err := Cmd("git", "log", "--pretty=fuller")
	if err != nil {
		return commits, err
	}

	commitArr := strings.Split(out, "commit")

	for _, v := range commitArr {
		arr := strings.Split(v, "\n")
		infos := make([]string, 0)
		for _, vv := range arr {
			str := strings.TrimSpace(vv)
			if len(str) > 0 {
				infos = append(infos, str)
			}
		}
		if len(infos) < 6 {
			continue
		}

		if len(infos) == 7 { // 包含merge，去除掉
			infos = append(infos[:1], infos[2:]...)
		}

		commit := &Commit{
			Hash:       getValue(infos[0]),
			Author:     getValue(infos[1]),
			AuthorDate: getDate(infos[2]),
			Commit:     getValue(infos[3]),
			CommitDate: getDate(infos[4]),
			Comment:    getValue(infos[5]),
		}
		commits = append(commits, commit)
	}

	return commits, nil
}

func EditAllCommit(name, email string) error {
	cmdStr := fmt.Sprintf(`
	CORRECT_NAME="%s"
	CORRECT_EMAIL="%s"
	
	export GIT_COMMITTER_NAME="$CORRECT_NAME"
	export GIT_COMMITTER_EMAIL="$CORRECT_EMAIL"
	export GIT_AUTHOR_NAME="$CORRECT_NAME"
	export GIT_AUTHOR_EMAIL="$CORRECT_EMAIL"
	`, name, email)

	out, err := Cmd("git", "filter-branch", "-f", "--env-filter", cmdStr)
	if err != nil {
		fmt.Printf("out: %v, error: %v\n", out, err)
		return err
	}

	return nil
}

func EditCommit(hash, date, name, email string) error {
	cmdStr := fmt.Sprintf(`
	CORRECT_DATE="%s"
	CORRECT_NAME="%s"
	CORRECT_EMAIL="%s"
	
	if [ "$GIT_COMMIT" = "%s" ]
	then
		export GIT_COMMITTER_DATE="$CORRECT_DATE"
		export GIT_COMMITTER_NAME="$CORRECT_NAME"
		export GIT_COMMITTER_EMAIL="$CORRECT_EMAIL"
		export GIT_AUTHOR_DATE="$CORRECT_DATE" 
		export GIT_AUTHOR_NAME="$CORRECT_NAME"
		export GIT_AUTHOR_EMAIL="$CORRECT_EMAIL"
	fi
	`, date, name, email, hash)

	out, err := Cmd("git", "filter-branch", "-f", "--env-filter", cmdStr)
	if err != nil {
		fmt.Printf("out: %v, error: %v\n", out, err)
		return err
	}

	return nil
}

func EditComment(hash, comment string) error {
	cmdStr := fmt.Sprintf(`
	if [ "$GIT_COMMIT" = "%s" ]
	then
		echo '%s'
	fi
	`, hash, comment)

	out, err := Cmd("git", "filter-branch", "-f", "--msg-filter", cmdStr)
	if err != nil {
		fmt.Printf("out: %v, error: %v\n", out, err)
		return err
	}

	return nil
}
