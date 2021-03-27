package main

import (
	"log"
	"path"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	files := []string{
		path.Join(curDir, "1.txt"),
		path.Join(curDir, "2.txt"),
		path.Join(curDir, "3.txt"),
		path.Join(curDir, "4.txt"),
	}
	//
	//for hash := range fileHashes(done, files...) {
	//	log.Printf("file hash: %q", hash)
	//
	//	// some action with hash
	//}

	for result := range fileHashesWithErrors(done, files...) {
		if result.err != nil {
			log.Printf("get hash error: %v", result.err)
			break
		}
		log.Printf("file hash: %q", result.hash)

		// some action with hash
	}
}
