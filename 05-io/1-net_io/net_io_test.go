package main

import (
	"bufio"
	"log"
	"net"
	"sync"
	"testing"
)

func Test_handleConn(t *testing.T) {
	// Имитация полнодуплексного
	// сетевого соединения.
	srv, cl := net.Pipe()

	// Мы должны дождаться
	// завершения всех потоков.
	var wg sync.WaitGroup
	wg.Add(1)

	// Обработчик подключения запускается
	// в отдельном потоке.
	go func() {
		// Обработчику передается один из
		// концов виртуальной трубы.
		handleConn(srv)
		srv.Close()
		wg.Done()
	}()

	// В основном потоке теста клиент отправляет
	// сообщения в трубу.
	// Эти сообщения сервер прочитает в соседнем потоке.
	_, err := cl.Write([]byte("time\n"))
	if err != nil {
		log.Fatal(err)
	}

	// Чтенеие ответа от сервера.
	reader := bufio.NewReader(cl)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Проверка ответа.
	if len(b) < 10 {
		t.Error("время не получено")
	}
	wg.Wait()
	cl.Close()
	t.Log(string(b))
}
