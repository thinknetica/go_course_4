package simple

import (
	"testing"
)

// Тест всегда имеет определённую сигнатуру: TestFuncName(t *testing.T) { ... }
func Test_sum(t *testing.T) {
	a, b := 3, 4 // тестовый пример
	want := 7    // заранее вычисленный результат

	got := sum(a, b) // вызов тестируемого кода

	if got != want { // сравнение результата с правильным значением
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}

	t.Log("OK")
}

// === RUN   Test_sum
// --- PASS: Test_sum (0.00s)
// PASS
// ok  	go-core/8-testing/simple	0.138s
