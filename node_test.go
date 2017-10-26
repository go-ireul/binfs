package binfs

import "testing"

func TestNodeBasic(t *testing.T) {
	n := node{}
	c := n.ensureChild("a")
	if c == nil {
		t.Fatal("child not created")
	}
	if c.name != "a" {
		t.Fatal("child name not set")
	}
	if len(c.path) != 1 || c.path[0] != "a" {
		t.Fatal("child path not set")
	}
	n.ensure("a", "b", "c")
	n.ensure("a", "c", "d")
	n.ensure("a", "c", "d", "e")
	c = n.find("a", "c", "d")
	if c == nil {
		t.Fatal("child not found")
	}
	if c.name != "d" {
		t.Fatal("child name not set")
	}
	if len(c.path) != 3 || c.path[0] != "a" || c.path[1] != "c" || c.path[2] != "d" {
		t.Fatal("child path not set")
	}
	if len(c.nodes) != 1 || c.nodes["e"] == nil || c.nodes["e"].name != "e" {
		t.Fatal("child nodes not set")
	}
}
