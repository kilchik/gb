package main

import (
	"log"
	"net/http"
	"time"
)

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(10*time.Second)
		log.Println("done")
	})
	http.ListenAndServe("localhost:7773", nil)
}
