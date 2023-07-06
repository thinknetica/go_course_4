package main

import (
	"fmt"
)

type computer struct {
	model string
	processor
}

type processor struct {
	model string
	cores int
}

func (c *computer) cpuinfo() string {
	return fmt.Sprintf("computer cpuinfo()")
}

func (p *processor) cpuinfo() string {
	return fmt.Sprintf("%s, %d ядер", p.model, p.cores)
}

func (c *computer) String() string {
	return fmt.Sprintf("Компьютер \"%s\".\nПроцессор \"%s\".\n", c.model, c.cpuinfo())
}

func main() {
	comp := &computer{
		model: "Компьютер игровой",
		processor: processor{
			model: "Байкал",
			cores: 8,
		},
	}

	fmt.Println(comp.cpuinfo())
}
