package main

import "log"

func computeHashesTerminatable(done <-chan interface{}, files <-chan string) <-chan interface{} {
	completed := make(chan interface{})

	go func() {
		defer log.Println("terminating child")
		defer close(completed)
		for {
			// делаем выбор между получением сообщения отмены и сообщением для выполнения работы
			select {
			case <-done:
				return

			case f := <-files:
				hash, err := compute(f)
				if err != nil {
					log.Printf("compute hash for %q: %v", f, err)
					continue
				}

				log.Printf("hash for file %q: %q", f, hash)
			}
		}
	}()

	return completed
}
