package main

import "net/http"

func main() {

	http.HandleFunc("/", index_main)
	http.ListenAndServe(":8080", nil)
	
}

func index_main(w http.ResponseWriter, r *http.Request)  {

	w.Write([]byte("<h1><center> Hello from Go! </h1></center>"))

}