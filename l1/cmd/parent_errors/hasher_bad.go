package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
)

func computeHashes(files <-chan string) <-chan interface{} {
	completed := make(chan interface{})

	go func() {
		defer close(completed)

		for file := range files {
			hash, err := compute(file)
			if err != nil {
				log.Printf("compute hash for %q: %v", file, err)
				continue
			}

			log.Printf("hash for file %q: %q", file, hash)
		}
	}()

	return completed
}


func compute(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", errors.Wrapf(err,"read file %q", path)
	}

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", errors.Wrap(err, "compute hash")
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

