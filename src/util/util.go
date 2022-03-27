package util

import "regexp"

func isGithub(input string) bool {
	res, _ := regexp.MatchString("github.com/*/*", input)
	return res
}
