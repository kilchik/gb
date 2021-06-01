package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const addr = "localhost:5301"

func main() {
	http.HandleFunc("/pay", func(writer http.ResponseWriter, request *http.Request) {
		bb, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Printf("read body: %v", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("paid for %q", bb)
	})
	log.Printf("listening %s...", addr)
	http.ListenAndServe(addr, nil)
}
