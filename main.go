package main

import "fmt"

func main() {
	cache := newCache(3)

	cache.put("A", "Alice")
	cache.put("B", "Brian")
	cache.put("C", "Charlie")

	printForward(cache.list.head)

	cache.get("A")
	printForward(cache.list.head)

	cache.put("B", "Byaku")
	printForward(cache.list.head)

	cache.put("D", "Diana")
	fmt.Println(cache.get("D"))
	printForward(cache.list.head)

	fmt.Println(cache.get("A"))
	fmt.Println(cache.get("C"))
}
