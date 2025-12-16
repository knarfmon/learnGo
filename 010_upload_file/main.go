package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

var s string

func foo(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		f, _, err := req.FormFile("usrfile")
		if err != nil {
			log.Println(err)
			http.Error(w, "Error uploading file", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		bs, err := io.ReadAll(f)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
}

//continued on
//https://www.youtube.com/watch?v=97ifBbsAii4
