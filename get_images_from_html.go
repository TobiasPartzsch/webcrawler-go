package main

import (
	"net/url"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	return resolveAttrBySelector(htmlBody, baseURL, "img[src]", "src")
}
