package main

import (
	log "github.com/kilchik/gb/lesson2/pkg/log"
	"io/ioutil"
	"os"
)

func main()  {
	log.Init(os.Stdout, os.Stderr, os.Stdout, ioutil.Discard)

	log.I.Printf("[host=%s] [uid=%d] file successfully uploaded", "srv42", 100500)
	log.W.Printf("[host=%s] [uid=%d] libjpeg: invalid format", "srv42", 200512)
	log.E.Printf("[host=%s] [uid=%d] file corrupted","srv42", 101345)
	log.I.Printf("[host=%s] storage space left: %d", "srv42", 1024)
}
