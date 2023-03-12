package trie

import (
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
)

// Algorithm Explanation: https://engineering.linecorp.com/ja/blog/simple-tries/

// Tree implements an immutable trie tree.
type Tree[K constraints.Ordered, V any] []node[K, V]

type node[K constraints.Ordered, V any] struct {
	value V
	next  int
	label K
	match bool
	leaf  bool
}

type kv[K constraints.Ordered, V any] struct {
	k []K
	v V
}

// New builds new Tree from keys and values. The len(keys) should equal to
// len(values).
func New[K constraints.Ordered, V any](keys [][]K, values []V) Tree[K, V] {
	if len(keys) != len(values) {
		panic("length mismatch of keys and values")
	}

	kvs := make([]kv[K, V], 0, len(keys))
	for i, k := range keys {
		kvs = append(kvs, kv[K, V]{k, values[i]})
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

	t := Tree[K, V]{node[K, V]{next: 1}}

	t = t.construct(kvs, 0, 0)
	return t
}

func (t Tree[K, V]) construct(kvs []kv[K, V], depth, current int) Tree[K, V] {
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
		t = append(t, node[K, V]{
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
func (t Tree[K, V]) Trace(path []K) Tree[K, V] {
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

// TraceOne is shorthand for Trace([]K{c}).
func (t Tree[K, V]) TraceOne(c K) Tree[K, V] {
	return t.Trace([]K{c})
}

// Terminal returns the value of the root of t. The second return value
// indicates whether the node has a value; if it is false, the first return
// value is zero value.
func (t Tree[K, V]) Terminal() (V, bool) {
	var zero V
	if len(t) == 0 {
		return zero, false
	}
	return t[0].value, t[0].match
}

// Predict returns the all values in the tree, t. The complexity is proportional
// to the number of nodes in t(it's not equal to len(t)).
func (t Tree[K, V]) Predict() []V {
	if len(t) == 0 || t[0].leaf {
		return nil
	}

	// Search linearly all of the child.
	var end int
	for !t[end].leaf {
		end = t.nextOf(t.nextOf(end)) - 1
	}

	var values []V
	for i := t.nextOf(0); i <= end; i++ {
		if t[i].match {
			values = append(values, t[i].value)
		}
	}
	return values
}

// Children returns the bytes of all direct children of the root of t. The result
// is sorted in ascending order.
func (t Tree[K, V]) Children() []K {
	if len(t) == 0 || t[0].leaf {
		return nil
	}

	var children []K
	for _, c := range t[t.nextOf(0):t.nextOf(t.nextOf(0))] {
		children = append(children, c.label)
	}
	return children
}

// nextOf returns the index of the next node of t[i].
func (t Tree[K, V]) nextOf(i int) int {
	return i + t[i].next
}
