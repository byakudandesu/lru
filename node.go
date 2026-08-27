package main

type node struct {
	prev *node
	next *node
	key  string
}

func NewNode(key string) *node {
	return &node{
		prev: nil,
		next: nil,
		key:  key,
	}
}

func GetKey(n *node) string {
	return n.key
}
