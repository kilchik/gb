package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main()  {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	func(ctx context.Context) {
		//http.Client{}
		//http.NewRequestWithContext()

		http.Get("localhost:7773")

	}(ctx)

	log.Println("stopped")
}
