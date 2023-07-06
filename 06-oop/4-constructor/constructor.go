package guitar

import (
	"fmt"
)

// Guitar - гитара.
type Guitar struct {
	manufacturer string
	model        string

	ch chan int
}

// New - конструктор.
func New(manufacturer, model string) *Guitar {
	g := Guitar{
		manufacturer: manufacturer, // имена совпадают
		model:        model,        // но это не проблема
	}
	g.ch = make(chan int)

	return &g // явно указываем, что возвращается ссылка
}

// Info возвращает сведения о гитаре.
func (g *Guitar) Info() string {
	return fmt.Sprintf("Гитара %s %s", g.manufacturer, g.model)
}
