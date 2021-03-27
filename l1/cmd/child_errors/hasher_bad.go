package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
)

func fileHashes(done <-chan interface{}, files ...string) <-chan string {
	hashes := make(chan string)

	go func() {
		defer close(hashes)

		for _, file := range files {
			hash, err := compute(file)
			if err != nil {
				// горутина пытается сообщить об ошибке, но к сожелению не может передать ошибку на уровень выше
				// сколько ошибок слишком много и нужно прервать выполнение горутины?
				// сколько можно чтобы продолжить делать запросы?
				log.Printf("compute hash for %q: %v", file, err)
				continue
			}

			select {
			case <-done:
				return
			case hashes <- hash:
			}
		}
	}()

	return hashes
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