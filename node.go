package binfs

type node struct {
	path  []string
	name  string
	nodes map[string]*node
	chunk *Chunk
}

func (n *node) ensureChild(name string) *node {
	if n.path == nil {
		n.path = []string{}
	}
	if n.nodes == nil {
		n.nodes = map[string]*node{}
	}
	c := n.nodes[name]
	if c == nil {
		c = &node{name: name, path: append(n.path, name)}
		n.nodes[name] = c
	}
	return c
}

func (n *node) ensure(name ...string) *node {
	t := n
	for _, v := range name {
		t = t.ensureChild(v)
	}
	return t
}

func (n *node) find(name ...string) *node {
	t := n
	for _, v := range name {
		if t != nil && t.nodes != nil {
			t = t.nodes[v]
		} else {
			return nil
		}
	}
	return t
}
