package main

import "fmt"

func main() {
	var i int = 10 // переменная
	var p *int     // указатель
	p = &i         // взятие адреса. указателю присваивается ссылка.
	var p1 *interface{}
	_ = p1
	fmt.Println(i)
	i++
	fmt.Println(i)
	*p++ // разыменование указателя
	fmt.Println(i)
	*(&i)++
	fmt.Println(i)
}

// Output: 	10
//			11
//			12
//			13
