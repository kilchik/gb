package main

import (
	"bufio"
	"fmt"
	"github.com/chzyer/readline"
	"io"
	"log"
	"os"
	"strings"
)

func main()  {
	db, err := os.OpenFile("todoer.db", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("cannot open db: %v", err)
	}
	defer db.Close()

	var items []string
	scanner := bufio.NewScanner(db)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	lineRdr, err := readline.NewEx(&readline.Config{
		Prompt:            "> ",
		HistoryFile:       "/tmp/todoer.tmp",
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
	})
	if err != nil {
		log.Fatalf("todoer: create line reader")
	}

	for {
		str, err := lineRdr.Readline()
		if err != nil {
			if err != readline.ErrInterrupt && err != io.EOF {
				log.Fatalf("read line: %v", err)
			}
			break
		}

		if str == "" {
			continue
		}
		tokens := strings.Split(str, " ")
		switch tokens[0] {
		case "add":
			if len(tokens) == 1 {
				continue
			}
			item := strings.Join(tokens[1:], " ")
			db.WriteString(item+"\n")
			items = append(items, item)
			fmt.Println(items)

		case "list":
			fmt.Println(strings.Join(items, "\n"))

		default:
			fmt.Println("unknown command")
		}
	}

}


