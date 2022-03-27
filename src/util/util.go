package util

import (
	"log"
	"regexp"
	"strings"
)

type GithubInfo struct {
	User string `json:"user"`
	Repo string `json:"repo"`
}

func IsGithub(input string) bool {
	res, _ := regexp.MatchString("github.com/.+/.+", input)
	return res
}

func ParseUserRepo(input string) GithubInfo {
	res := GetStringInBetween(input, "github.com/", "SDUIE$FHWUIFNKWF$")
	user := strings.Split(res, "/")[0]
	repo := strings.Split(res, "/")[1]
	return GithubInfo{
		User: user,
		Repo: repo,
	}
}

func GetStringInBetween(str, start, end string) string {
	s := strings.Index(str, start)
	if s == -1 {
		log.Fatal("start of string is failing")
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		log.Fatal("end of string is failing")
	}
	return str[s : s+e]
}
