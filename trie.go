package trie

import (
	"fmt"
	"sort"
)

// Algorithm Explanation: https://engineering.linecorp.com/ja/blog/simple-tries/

// Tree implements an immutable trie tree.
type Tree []node

type node struct {
	value interface{}
	next  int
	label byte
	match bool
	leaf  bool
}

type kv struct {
	k []byte
	v interface{}
}

// New builds new Tree from keys and values. The len(keys) should equal to
// len(values).
func New(keys [][]byte, values []interface{}) Tree {
	if len(keys) != len(values) {
		panic("length mismatch of keys and values")
	}

	kvs := make([]kv, 0, len(keys))
	for i, k := range keys {
		kvs = append(kvs, kv{k, values[i]})
	}
	sort.Slice(kvs, func(i, j int) bool {
		a, b := kvs[i].k, kvs[j].k
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] == b[i] {
				continue
			}
			return a[i] < b[i]
		}
		if len(a) == len(b) {
			panic(fmt.Sprintf("2 same key is passed: %v", kvs[i].k))
		}
		return len(a) < len(b)
	})

	t := Tree{node{next: 1}}

	t = t.construct(kvs, 0, 0)
	return t
}

func (t Tree) construct(kvs []kv, depth, current int) Tree {
	if depth == len(kvs[0].k) {
		t[current].match = true
		t[current].value = kvs[0].v
		kvs = kvs[1:]
		if len(kvs) == 0 {
			t[current].leaf = true
			return t
		}
	}

	p := []int{0}
	for i := 0; i < len(kvs); {
		t = append(t, node{
			label: kvs[i].k[depth],
		})
		for c := kvs[i].k[depth]; i < len(kvs) && kvs[i].k[depth] == c; i++ {
		}
		p = append(p, i)
	}

	for i := 0; i < len(p)-1; i++ {
		t[t.nextOf(current)+i].next = len(t) - t.nextOf(current) - i
		t = t.construct(kvs[p[i]:p[i+1]], depth+1, t.nextOf(current)+i)
	}
	return t
}

// Trace returns the subtree of t whose root is the node traced from the root
// of t by path. It doesn't modify t itself, but returns the subtree.
func (t Tree) Trace(path []byte) Tree {
	if t == nil {
		return nil
	}

	var u int
	for _, c := range path {
		if t[u].leaf {
			return nil
		}
		u = t.nextOf(u)
		v := t.nextOf(u)
		if v-u > 40 {
			// Binary Search
			u += sort.Search(v-u, func(m int) bool {
				return t[u+m].label >= c
			})
		} else {
			// Linear Search
			for ; u != v-1 && t[u].label < c; u++ {
			}
		}
		if u > v || t[u].label != c {
			return nil
		}
	}
	return t[u:]
}

// TraceByte is shorthand for Trace([]byte{c}).
func (t Tree) TraceByte(c byte) Tree {
	return t.Trace([]byte{c})
}

// Terminal returns the value of the root of t. The second return value
// indicates whether the node has a value; if it is false, the first return
// value is nil. It returns nil also when the t is nil.
func (t Tree) Terminal() (interface{}, bool) {
	if len(t) == 0 {
		return nil, false
	}
	return t[0].value, t[0].match
}

// Predict returns the all values in the tree, t. The complexity is proportional
// to the number of nodes in t(it's not equal to len(t)).
func (t Tree) Predict() []interface{} {
	if len(t) == 0 || t[0].leaf {
		return nil
	}

	// Search linearly all of the child.
	var end int
	for !t[end].leaf {
		end = t.nextOf(t.nextOf(end)) - 1
	}

	var values []interface{}
	for i := t.nextOf(0); i <= end; i++ {
		if t[i].match {
			values = append(values, t[i].value)
		}
	}
	return values
}

// Children returns the bytes of all direct children of the root of t. The result
// is sorted in ascending order.
func (t Tree) Children() []byte {
	if len(t) == 0 || t[0].leaf {
		return nil
	}

	var children []byte
	for _, c := range t[t.nextOf(0):t.nextOf(t.nextOf(0))] {
		children = append(children, c.label)
	}
	return children
}

// nextOf returns the index of the next node of t[i].
func (t Tree) nextOf(i int) int {
	return i + t[i].next
}
