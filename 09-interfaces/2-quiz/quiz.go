package main

import (
	"fmt"
)

func main() {
	var iface any        // 0
	var f func()         // 1
	var m map[string]int // 2
	var p *int = nil
	var iface2 any = p // 3
	cycle(iface, f, m, iface2)
}

func cycle(ifaces ...any) {
	for i, iface := range ifaces {
		if iface == nil {
			fmt.Println(i)
		}
	}
}
