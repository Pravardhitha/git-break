package util

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// github api

func GET(user, repo string) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := fmt.Sprintf("https://github.com/%s/%s/", user, repo)
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
	fmt.Println(res)

}
