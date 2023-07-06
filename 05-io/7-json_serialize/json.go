package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	Lon float64 `json:"1"`
	Lat float64 `json:"2"`
}

func main() {
	var points []Point
	for i := 0; i < 100; i++ {
		points = append(points, Point{Lon: float64(i), Lat: float64(i)})
	}

	b, _ := json.MarshalIndent(points, "", "  ")
	fmt.Println(string(b))
}
