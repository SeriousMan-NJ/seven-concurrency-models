package main

import (
	"testing" // 테스트 코드는 항상 testing 패키지를 가져옴
)

func TestReadAll(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	r := ReadAll(ch)

	for i, v := range []int{0, 1, 2, 3, 4} {
		if v != r[i] {
			t.Error()
		}
	}
}

func TestWriteAll(t *testing.T) {
	ch := make(chan int)
	l := []int{1, 2, 3, 4, 5}

	go func() {
		WriteAll(ch, l)
	}()

	r := ReadAll(ch)

	for i, v := range l {
		if v != r[i] {
			t.Error()
		}
	}
}
