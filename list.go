package main

import "fmt"

type linkedList struct {
	head *node
	tail *node
}

func printForward(start *node) {
	keys := []string{}
	current := start
	for current != nil {
		keys = append(keys, current.key)
		current = current.next
	}
	fmt.Println(keys)
}

func printBackward(start *node) {
	keys := []string{}
	current := start
	for current != nil {
		keys = append(keys, current.key)
		current = current.prev
	}
	fmt.Println(keys)
}

func (list *linkedList) detach(n *node) {
	if n != list.head {
		n.prev.next = n.next
	}
	if n != list.tail {
		n.next.prev = n.prev
	}
	if n == list.head {
		list.head = n.next
	}
	if n == list.tail {
		list.tail = n.prev
	}
	n.next = nil
	n.prev = nil
}

func (list *linkedList) insertAtHead(n *node) {
	if list.head == nil {
		list.head = n
		list.tail = n
		return
	}
	list.head.prev = n
	n.next = list.head
	list.head = n
}
