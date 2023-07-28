package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

// Сетевой адрес.
//
// Служба будет слушать запросы на всех IP-адресах
// компьютера на порту 12345.
// Нпример, 127.0.0.1:12345
const addr = "0.0.0.0:12345"

// Протокол сетевой службы.
const proto = "tcp4"

func main() {
	// Запуск сетевой службы про протоколу TCP
	// на порту 12345.
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	// Подключения обрабатываются в бесконечном цикле.
	// Иначе после обслуживания первого подключения сервер
	//завершит работу.
	for {
		// Принимаем подключение.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Вызов обработчика подключения.
		go handleConn(conn)
	}
}

// Обработчик. Вызывается для каждого соединения.
func handleConn(conn net.Conn) {
	// Чтение сообщения от клиента.
	reader := bufio.NewReader(conn)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Println(err)
		return
	}

	// Удаление символов конца строки.
	msg := strings.TrimSuffix(string(b), "\n")
	msg = strings.TrimSuffix(msg, "\r")

	// Если получили "time" - пишем время в соединение.
	if msg == "time" {
		conn.Write([]byte(time.Now().String() + "\n"))
	}

	// Закрытие соединения.
	conn.Close()
}
