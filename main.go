package main

import "fmt"

func main() {
	A := NewNode("A")
	B := NewNode("B")

	A.next = B
	B.prev = A

	fmt.Println(GetKey(A))
	fmt.Println(GetKey(B))
	C := NewNode("C")

	B.next = C
	C.prev = B

	list := &linkedList{}

	list.head = A
	list.tail = C
	printForward(A)
	printBackward(C)

	list.detach(A)
	printForward(list.head)
	printBackward(list.tail)
}
