package main

import (
	"log"
	"time"
)

func main() {
	//var files chan string
	//
	//completed := computeHashes(files)
	//
	//// тут может быть ещё какая-либо логика работы приложения
	//<-completed
	//log.Println("done.")

	var files chan string

	done := make(chan interface{})
	terminated := computeHashesTerminatable(done, files)

	go func() {
		// отменяем горутину внутри computeHashesTerminatable через одну секунду
		time.Sleep(1 * time.Second)
		log.Println("canceling computeHashesTerminatable goroutine")
		close(done)
	}()

	<-terminated

	log.Println("done")
}
