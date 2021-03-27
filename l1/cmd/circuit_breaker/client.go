package main

import (
	"log"
	"net"
	"time"

	"bufio"
	"fmt"
)

const addr = "127.0.0.1:5433"

func main() {
	timeout := 1 * time.Second

	orders := []string{
		"123432-234",
		"123432-234",
		"123432-234",
		"853432-332",
		"853432-332",
		"853432-332",
		"254432-341",
		"254432-341",
		"254432-341",
		"254432-341",
		"853432-332",
		"853432-332",
		"853432-332",
		"254432-341",
		"123432-234",
		"123432-234",
		"853432-332",
		"123432-234",
	}

	for _, order := range orders {
		// паттерн circuit breaker в простейшей реализации
		// в реальности за функцией может скрываться объёмная и сложная логика
		// лишнее выполнение которой мы и хотим прервать
		conn, ok := checkCircuitOpen(addr, timeout)
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}
		defer conn.Close()

		// если сервер стал доступен, возвращаем логику работы с ним
		pay(order, conn)
		time.Sleep(1 * time.Second)
	}
}

func checkCircuitOpen(address string, timeout time.Duration) (net.Conn, bool) {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		log.Printf("can't connect to server, error: %s", err)
		return nil, false
	}

	return conn, true
}

func pay(order string, conn net.Conn) error {
	fmt.Fprintf(conn, "please, process order %q\n", order)
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}
	log.Printf("Message from server: %q", message)

	return nil
}
