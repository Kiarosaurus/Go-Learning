package main

import "fmt"

// List representa una lista enlazada individualmente que contiene
// valores de cualquier tipo.
type List[T any] struct {
	next *List[T]
	val  T
}

// PushBack inserta 'val' al final de la lista con head 'l'.
func (l *List[T]) PushBack (val T) bool {
	if l == nil {	// Lista nil.
		return false
	}
		
	aux := l
	for aux.next != nil {
		aux = aux.next
	}
	aux.next = &List[T]{val: val}
	return true
}


// Print omitiendo nodo centinela.
func (l *List[T]) PrettyPrint() []T {
	var out []T
	for n := l.next; n != nil; n = n.next {
		out = append(out, n.val)
	}
	return out
}

func main() {
	var head List[int]	// cabeza centinela
	head.PushBack(11)
	head.PushBack(22)
	head.PushBack(33)
	fmt.Println(head.PrettyPrint())
}
