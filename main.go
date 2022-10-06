package main

import (
	"fmt"
	"net/http"
)

type server struct {
	mux *http.ServeMux
}

func (s *server) routes() {
	s.mux.HandleFunc("/foo/bar", foohandle)
	s.mux.HandleFunc("/baz", bazhandle)
}

func NewServer() server {
	s := server{
		mux: http.NewServeMux(),
	}

	s.routes()
	return s
}

func bazhandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("req: ", req)
	res.Write([]byte("baz!"))
}

func foohandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("req: ", req)
	res.Write([]byte("foo!"))
}

func main() {
	setup()

	//for wasm prevent from closing
	<-make(chan bool)
}
