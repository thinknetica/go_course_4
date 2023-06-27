// Пример реализации структуры данных "Двоичное дерево поиска".
// Пример для простоты приведён для дерева,
// содержащего в качетве значений целые числа.
// Можно по аналогии со стандартной библиотекой с помощью интерфейса
// обобщить на произвольные типы данных.
// Для этого потребуется контракт на функцию сравнения элементов.
package main

import "fmt"

// Двоичное дерево поиска
type Tree struct {
	root *Element
}

// Лист дерева.
type Element struct {
	left, right *Element
	Value       int
}

// Insert вставляет элемент в дерево.
func (t *Tree) Insert(x int) {
	e := &Element{Value: x}
	if t.root == nil {
		t.root = e
		return
	}
	insert(t.root, e)
}

// inset рекурсивно вставляет элемент в нужный уровень дерева.
func insert(node, new *Element) {
	if new.Value < node.Value {
		if node.left == nil {
			node.left = new
			return
		}
		insert(node.left, new)
	}
	if new.Value >= node.Value {
		if node.right == nil {
			node.right = new
			return
		}
		insert(node.right, new)
	}
}

// Search ищет значение в дереве.
func (t *Tree) Search(x int) bool {
	return search(t.root, x)
}

func search(el *Element, x int) bool {
	if el == nil {
		return false
	}
	if el.Value == x {
		return true
	}
	if el.Value < x {
		return search(el.right, x)
	}
	return search(el.left, x)
}

// String реализует интерфейс Stringer для дерева.
func (t Tree) String() string {
	return prettyPrint(t.root, 0)
}

// prettyPrint печатает дерево в виде дерева :)
func prettyPrint(e *Element, spaces int) (res string) {
	if e == nil {
		return res
	}
	spaces++
	res += prettyPrint(e.right, spaces)
	for i := 0; i < spaces; i++ {
		res += "\t"
	}
	res += fmt.Sprintf("%d\n", e.Value)
	res += prettyPrint(e.left, spaces)
	return res
}

func initTree() *Tree {
	var t Tree
	t.Insert(10)
	t.Insert(15)
	t.Insert(15)
	t.Insert(15)
	t.Insert(15)
	t.Insert(25)
	t.Insert(30)
	t.Insert(35)
	t.Insert(5)
	t.Insert(20)
	t.Insert(1)
	t.Insert(2)
	t.Insert(6)
	return &t
}

func main() {
	t := initTree()
	fmt.Println(t)
}

// 	Output:
// 					35
// 				30
// 			25
// 		20
// 			15
// 	10
//			6
// 		5
// 				2
// 			1
