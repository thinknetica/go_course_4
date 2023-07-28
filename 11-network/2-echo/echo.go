package main

// Эхо-сервер. Без цензуры.

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// обработчик подключения
func handler(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("Connection Closed")

	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		msg = append(msg, '\n')

		_, err = conn.Write(msg)
		if err != nil {
			return
		}
	}
}

func main() {
	// регистрация сетевой службы
	listener, err := net.Listen("tcp4", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Сервер слушает на порту: 12345")

	// цикл обработки клиентских подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn)
	}
}
