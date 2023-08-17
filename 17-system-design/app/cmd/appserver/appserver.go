package main

import "go-core-4/17-system-design/app/internal/server"

func main() {
	s, _ := server.New()
	s.Run()
}
