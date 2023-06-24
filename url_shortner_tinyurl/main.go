package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var longURL string
    fmt.Scanf("Enter the link to shorten: %s", &longURL)

	shortURL, err := shortenURL(longURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Shortened URL:", shortURL)
}

func shortenURL(longURL string) (string, error) {
	apiURL := "http://tinyurl.com/api-create.php?url=" + longURL

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	shortURL := strings.TrimSpace(string(body))
	return shortURL, nil
}
