package main

type cache struct {
	list     linkedList
	hashmap  map[string]*node
	capacity int
}

func newCache(cap int) *cache {
	return &cache{
		list:     linkedList{},
		hashmap:  make(map[string]*node),
		capacity: cap,
	}
}

func (c *cache) add(n *node) {
	c.list.insertAtHead(n)
	c.hashmap[n.key] = n
}

func (c *cache) get(k string) (string, bool) {
	node, ok := c.hashmap[k]
	if ok {
		if c.list.head != node {
			c.list.detach(node)
			c.list.insertAtHead(node)
		}
		return node.value, true
	}
	return "", false
}

func (c *cache) put(k, v string) {
	node, ok := c.hashmap[k]
	if ok {
		if c.list.head != node {
			c.list.detach(node)
			c.list.insertAtHead(node)
		}
		node.value = v
		return
	}
	n := newNode(k, v)
	c.add(n)

	if len(c.hashmap) > c.capacity {
		evicted := c.list.tail
		c.list.detach(evicted)
		delete(c.hashmap, evicted.key)
	}
}
