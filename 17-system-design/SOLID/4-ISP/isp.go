package isp

// Пакет демонстрирует принцип разделения интерфейсов.

// Интерфейс, включающий в себя все возможные методы,
// которые используются в пакете.
type BigInterface interface {
	Dance()
	Sing()
	PlayGuitar()
}

// Минимальный интерфейс, требуемый
// для функции Dance.
type Dancer interface {
	Dance()
}

// Нарушение ISP, поскольку от аргумента требуется
// реализация методов, которые не используются.
func BadDance(dancer BigInterface) {
	dancer.Dance()
}

// Соблюдение ISP за счёт разделения интерфейсов
// для разных задач.
func GoodDance(dancer Dancer) {
	dancer.Dance()
}
