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

type SpecificLicense struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	SPDX_ID string `json:"spdx_id"`
	Url     string `json:"url"`
	NODE_ID string `json:"node_id"`
}

type LicenseReturn struct {
	DoesRet bool
	Value   SpecificLicense
}

type Verification struct {
	Verified  bool   `json:"verified"`
	Reason    string `json:"reason"`
	Signature string `json:"signature"`
	Payload   string `json:"payload"`
}

// Branches: checks Branch protections
// Implicitly this has at least one branch so doesn't error not found it
func Branches(user, repo string) []BranchSec {
	var jsonData []BranchSec
	url := fmt.Sprintf("repos/%s/%s/branches", user, repo)
	retData := GET(url)

	json.Unmarshal([]byte(retData), &jsonData)
	return jsonData
}

func License(user, repo string) LicenseReturn {
	var jsonData map[string]interface{}
	var ret LicenseReturn
	url := fmt.Sprintf("repos/%s/%s/license", user, repo)
	retData := GET(url)

	json.Unmarshal([]byte(retData), &jsonData)
	if val, ok := jsonData["license"]; ok {
		var lic SpecificLicense
		a, _ := json.Marshal(val)
		json.Unmarshal([]byte(a), &lic)

		ret.DoesRet = ok
		ret.Value = lic
	} else {
		ret.DoesRet = ok
		ret.Value = SpecificLicense{}
	}
	return ret
}

// Implicitly this has at least one commit so doesn't error not found it
func Commits(user, repo string) {
	var jsonData map[string]interface{}
	url := fmt.Sprintf("repos/%s/%s/commits", user, repo)
	retData := GET(url)
	json.Unmarshal([]byte(retData), &jsonData)

	if commitVal, ok := jsonData["commit"]; ok {
		var insideCommit map[string]interface{}
		jsonCommit, _ := json.Marshal(commitVal)
		json.Unmarshal([]byte(jsonCommit), &insideCommit)
		// TODO
	}
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
