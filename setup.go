//go:build !js
// +build !js

package main

import (
	"fmt"
	"net/http"
)

func setup() {
	s := NewServer()
	err := http.ListenAndServe(":9001", s.mux)
	if err != nil {
		fmt.Println("error: ", err)
	}
}
