//go:build js
// +build js

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

type ResponseWriter struct {
	Body *bytes.Buffer
}

func (r ResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (r ResponseWriter) Write(buf []byte) (int, error) {
	if r.Body != nil {
		r.Body.Write(buf)
	}
	return len(buf), nil
}

func (r ResponseWriter) WriteHeader(statusCode int) {}

func createWriter() ResponseWriter {
	return ResponseWriter{
		Body: new(bytes.Buffer),
	}
}

func parsePath(this js.Value, args []js.Value) interface{} {
	fmt.Println("args: ", args)
	server := NewServer()
	path := args[0].String()
	req, err := http.NewRequest(http.MethodGet, path, nil)

	if err != nil {
		fmt.Println("request error: ", err)
	}

	h, p := server.mux.Handler(req)

	fmt.Println("pattern: ", p)

	w := createWriter()

	h.ServeHTTP(w, req)

	b, err := ioutil.ReadAll(w.Body)

	if err != nil {
		fmt.Println("ioutil error: ", err)
	}

	return string(b)
}

func setup() {
	js.Global().Set("parsePath", js.FuncOf(parsePath))
}
