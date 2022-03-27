package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type BranchSec struct {
	Name      string       `json:"name"`
	Commit    CommitBranch `json:"commit"`
	Protected bool         `json:"protected"`
}

type CommitBranch struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// Branches: checks Branch protections
func Branches(user, repo string) []BranchSec {
	var jsonData []BranchSec
	url := fmt.Sprintf("repos/%s/%s/branches", user, repo)
	retData := GET(url)

	json.Unmarshal([]byte(retData), &jsonData)
	return jsonData
}

// GET: Generic GET request to send to generic github api
func GET(endpoint string) string {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := fmt.Sprintf("https://api.github.com/%s", endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
