package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	tenSecondDuration = 10 * time.Second
)

func Racer(a, b string) (winder string, error error) {
	return ConfigurableRacer(a, b, tenSecondDuration)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// Given two URLS and a timeout duration, returns the URL that responds first,
	// or an error if neither URL responds within the timeout duration.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
