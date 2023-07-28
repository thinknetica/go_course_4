package main

// Эхо-сервер. Без цензуры.

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// обработчик подключения
func handler(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("Conn closed")

	conn.SetDeadline(time.Now().Add(time.Second * 5))

	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		fmt.Println(string(msg))
		_, err = conn.Write(msg)
		if err != nil {
			return
		}

		conn.SetDeadline(time.Now().Add(time.Second * 5))
	}
}

func main() {
	// регистрация сетевой службы
	listener, err := net.Listen("tcp4", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	// цикл обработки клиентских подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Conn established")

		go handler(conn)
	}
}
