package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/unboxed/loganthecat/amazon"
	"github.com/unboxed/loganthecat/github"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Commit Received")

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	payload, err := github.ParseCommit(body)

	if err != nil {
		return
	}
	log.Println(payload)

	go amazon.Run()
}

func main() {
	http.HandleFunc("/hooks/commit", handler)
	http.ListenAndServe(":3001", nil)
}
