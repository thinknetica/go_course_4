package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // буфер для os.Stdin
	for {
		fmt.Print("NewReader-> ")
		text, _ := reader.ReadString('\n')    // чтение строки (до символа перевода)
		text = strings.TrimSuffix(text, "\n") // удаление перевода строки
		text = strings.TrimSuffix(text, "\r") // удаление перевода строки для Windows
		if text == "exit" {
			break
		}
		fmt.Println("echo:", text)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("NewScanner-> ")
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			break
		}
		fmt.Println(scanner.Text())
		fmt.Print("NewScanner-> ")
	}

	for {
		fmt.Print("fmt.Scanln-> ")
		var s string
		fmt.Scanln(&s)
		if s == "exit" {
			break
		}
		fmt.Println(s)
	}
}
