package main

import "fmt"

const greeting = "Hello, world!"

func main() {
	printMsg(greeting)
}

func printMsg(msg string) {
	fmt.Println(greeting)
}

// Output: Hello, world!
