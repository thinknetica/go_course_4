package main

// Сервер текущего времени в соответсвии с RFC 867.

import (
	"io"
	"log"
	"net"
	"time"
)

// обработчик подключения
func handler(conn io.ReadWriteCloser) {
	daytime := time.Now().Format(time.RubyDate)
	conn.Write([]byte(daytime))
	conn.Close()
}

func main() {
	// регистрация сетевой службы на всех сетевых интерфейсах на порту 13
	listener, err := net.Listen("tcp4", "0.0.0.0:13")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// цикл обработки клиентских подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		handler(conn)
	}
}
