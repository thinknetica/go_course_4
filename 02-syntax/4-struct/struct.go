package main

import "fmt"

type point struct {
	lat float64
	lon float64
}

func (p point) PrintPoint() {
	fmt.Printf("lat:lon - %f, %f\n", p.lat, p.lon)
}

type object struct {
	name  string
	point // встраивание, обращение по имени типа
}

func main() {
	o := object{}
	o.name = "Object"
	o.lat = 10
	o.lon = 20
	o.PrintPoint()
}
