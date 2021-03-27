package main

import (
	"bufio"
	"github.com/pkg/errors"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:5433")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("server is running...")

	for {
		if err := processPayment(listener); err != nil {
			log.Printf("client connect error: %s", err)
		}
	}
}

func processPayment(listener net.Listener) error {
	conn, err := listener.Accept()
	if err != nil {
		return errors.Wrap(err, "accept new connection")
	}
	defer conn.Close()

	log.Print("received connection")
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return errors.Wrap(err, "get message from client")
	}

	log.Printf("new message received: %q", message)

	// Выполняем обработку платежа

	// Возвращаем клиенту обработанное сообщение
	if _, err := conn.Write([]byte("alrighty\n")); err != nil {
		return errors.Wrap(err, "write response to client")
	}

	return nil
}
