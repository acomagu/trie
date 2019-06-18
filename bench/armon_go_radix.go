package bench

import (
	"github.com/armon/go-radix"
)

type armonGoRadix struct {
	tree *radix.Tree
}

func (t *armonGoRadix) Name() string {
	return "github.com/armon/go-radix"
}

func (t *armonGoRadix) Build(keys [][]byte, values []interface{}) {
	t.tree = radix.New()
	for i := range keys {
		t.tree.Insert(string(keys[i]), values[i])
	}
}

func (t *armonGoRadix) Get(s []byte) (interface{}, bool) {
	return t.tree.Get(string(s))
}

func (t *armonGoRadix) LongestPrefix(s []byte) (interface{}, bool) {
	_, v, match := t.tree.LongestPrefix(string(s))
	return v, match
}
