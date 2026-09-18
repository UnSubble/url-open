package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) string {
	if strings.Contains(rawURL, "://") {
		return rawURL
	}

	return "http://" + rawURL
}

func extractDomain(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return u.Host, nil
}
