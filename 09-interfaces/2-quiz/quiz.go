package main

import (
	"fmt"
)

type myMap map[int]int

func main() {
	var iface interface{} // 0
	var f func()          // 1
	var m myMap = nil     // 2

	cycle(iface, f, m)
}

func cycle(ifaces ...interface{}) {
	for i, iface := range ifaces {
		if iface == nil {
			fmt.Println(i)
		}
	}
}
