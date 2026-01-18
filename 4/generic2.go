package main

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	Data T
	next *node[T]
}

type list[T any] struct {
	start *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{Data: data,
		next: nil,
	}
	if l.start == nil {
		l.start = &n
		return
	}
	if l.start.next == nil {
		l.start.next = &n
		return
	}
	temp := l.start
	l.start = l.start.next
	l.add(n.Data)
	l.start = temp

}

type treeLast[T any] []T

func (t treeLast[T]) replaceLast(elem T) ([]T, error) {
	if len(t) == 0 {
		return t, errors.New("this is empty")
	}
	t[len(t)-1] = elem
	return t, nil

}

func main() {
	ts := treeLast[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 99}
	tt := treeLast[string]{"banana", "apple", "peach", "date"}
	fmt.Println(ts, "\\", tt, 123)
	ts.replaceLast(100999)
	tt.replaceLast("Ali")
	fmt.Println(ts, "\\", tt, 123)

}
