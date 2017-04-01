package main

import (
	"net/http"
	"golang.org/x/net/http2"
)

func main() {

	var srv http.Server
	srv.Addr = ":8080"
	//Enable http2
	http2.ConfigureServer(&srv, nil)

	http.HandleFunc("/", index_main)

	srv.ListenAndServeTLS("certs/localhost.cert", "certs/localhost.key")
	
}

func index_main(w http.ResponseWriter, r *http.Request)  {

	w.Write([]byte("<h1><center> Hello from Go! </h1></center>"))

}