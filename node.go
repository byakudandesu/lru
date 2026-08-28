package main

type node struct {
	prev  *node
	next  *node
	key   string
	value string
}

func newNode(key, value string) *node {
	return &node{
		prev:  nil,
		next:  nil,
		key:   key,
		value: value,
	}
}

func getKey(n *node) string {
	return n.key
}

func getValue(n *node) string {
	return n.value
}
