package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main()  {



	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	hc := http.Client{}

	func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return

		default:
			log.Println("start http.Get")
			req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:7773/", nil)
			hc.Do(req)
		}

	}(ctx)

	log.Println("stopped")
}







//hc := http.Client{
//	Timeout: 2 * time.Second,
//}

//req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:7773/", nil)