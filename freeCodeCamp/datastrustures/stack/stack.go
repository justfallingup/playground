package stack

import (
	"fmt"
)

// Stack - стэк на основе двусвязного списка.
type Stack struct {
	root *Elem
}

// Elem - элемент стэка.
type Elem struct {
	Val interface{}
	next, prev *Elem
}

// New создаёт стэк и возвращает указатель на него.
func New() *Stack {
	var s Stack
	s.root = &Elem{}
	s.root.prev = s.root
	s.root.next = s.root
	return &s
}

// Empty проверяет наличие непустых элементов в стэке
func (s *Stack) Empty() bool {
	if s.root.next == s.root {
		return true
	}
	return false
}

// Push вставляет элемент с Val наверх стэка.
func (s *Stack) Push(val interface{}) *Stack {
	e := Elem{Val: val}
	e.prev = s.root
	e.next = s.root.next
	s.root.next = &e
	if e.next != s.root {
		e.next.prev = &e
	}
	return s
}

// Pop удаляет верхний элемент стэка, возвращает Val этого элемента
func (s *Stack) Pop() interface{} {
	el := s.root.next
	if el == s.root {
		return nil
	}
	s.root.next = el.next
	el.next.prev = s.root
	val := el.Val
	return val
}

// String реализует интерфейс fmt.Stringer представляя стэк в виде строки.
func (s *Stack) String() string {
	el := s.root.next
	var st string
	for el != s.root {
		st += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	if len(st) > 0 {
		st = st[:len(st)-1]
	}
	return st
}

// Reverse разворачивает стэк.
func (s *Stack) Reverse() *Stack {
	el := s.root.next
	if el == s.root {
		return s
	}
	for el != s.root {
		el.prev, el.next = el.next, el.prev
		if el.prev == s.root {
			break
		}
		el = el.prev
	}
	s.root.next = el
	return s
}