package main

type result struct {
	err error
	hash string
}

func fileHashesWithErrors(done <-chan interface{}, files ...string) <-chan result {
	results := make(chan result)

	go func() {
		defer close(results)

		for _, file := range files {
			hash, err := compute(file)
			res := result{
				err:  err,
				hash: hash,
			}

			select {
			case <-done:
				return
			case results <- res:
			}
		}
	}()

	return results
}
