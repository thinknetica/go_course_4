package interview

import (
	"fmt"
	"sync"
)

func TypicalErrorsNoSync() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func TypicalErrorsIndex() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}

func Slices() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[1:3]
	fmt.Println("До изменения S2")
	fmt.Println(s1)
	fmt.Println(s2)

	s2[1] = 999
	fmt.Println("После изменения S2")
	fmt.Println(s1)
	fmt.Println(s2)

	s2 = append(s2, 5, 6, 7, 8, 9)
	fmt.Println("После добавления к S2")
	fmt.Println(s1)
	fmt.Println(s2)
}

func Strings() {
	s := "Привет!"

	for index, rune := range s {
		fmt.Println(index, rune)
	}
}
