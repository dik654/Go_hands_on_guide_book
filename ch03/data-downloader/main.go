package main

import "net/http"

func fetchRemoteResource(url string) ([]byte, error) {
	r, err := http.Get(url)
}