package main

import "fmt"

func main() {
	cache := newCache(3)
	A := newNode("A", "Alice")
	B := newNode("B", "Brian")
	C := newNode("C", "Charlie")

	cache.add(A)
	cache.add(B)
	cache.add(C)

	printForward(cache.list.head)

	cache.get("A")
	printForward(cache.list.head)

	cache.put("B", "Byaku")
	printForward(cache.list.head)

	cache.put("D", "Diana")
	fmt.Println(cache.get("D"))
	printForward(cache.list.head)
}
