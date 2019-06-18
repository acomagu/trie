package bench

import (
	"github.com/kkdai/radix"
)

type kkdaiRadix struct {
	tree *radix.RadixTree
}

func (t *kkdaiRadix) Name() string {
	return "github.com/kkdai/radix"
}

func (t *kkdaiRadix) Build(keys [][]byte, values []interface{}) {
	t.tree = radix.NewRadixTree()
	for i := range keys {
		t.tree.Insert(string(keys[i]), values[i])
	}
}

func (t *kkdaiRadix) Get(s []byte) (interface{}, bool) {
	return t.tree.Lookup(string(s))
}
