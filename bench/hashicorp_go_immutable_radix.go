package bench

import (
	iradix "github.com/hashicorp/go-immutable-radix"
)

type hashicorpGoImmutableRadix struct {
	tree *iradix.Tree
}

func (t *hashicorpGoImmutableRadix) Name() string {
	return "github.com/hashicorp/go-immutable-radix"
}

func (t *hashicorpGoImmutableRadix) Build(keys [][]byte, values []interface{}) {
	t.tree = iradix.New()
	for i := range keys {
		t.tree, _, _ = t.tree.Insert(keys[i], values[i])
	}
}

func (t *hashicorpGoImmutableRadix) Get(s []byte) (interface{}, bool) {
	return t.tree.Get(s)
}

func (t *hashicorpGoImmutableRadix) LongestPrefix(s []byte) (interface{}, bool) {
	root := t.tree.Root()
	_, v, matched := root.LongestPrefix(s)
	return v, matched
}
