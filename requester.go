package main

import (
	"net/http"
	"strings"
)

func IsOK(url string) bool {
	url = FormatURL(url)

	resp, err := http.Get(url)
	if err != nil {
		return false
	}

	return resp.StatusCode == http.StatusOK
}

func FormatURL(url string) string {
	newURL := strings.ReplaceAll(url, "https:/", "https://")
	newURL = strings.ReplaceAll(newURL, "http:/", "http://")

	if !strings.Contains(newURL, "http") {
		newURL = "http://" + newURL
	}

	return newURL
}
