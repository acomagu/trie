package bench

import (
	"strings"

	"github.com/sauerbraten/radix"
)

type sauerbratenRadix struct {
	tree *radix.Radix
}

func (t *sauerbratenRadix) Name() string {
	return "github.com/sauerbraten/radix"
}

func (t *sauerbratenRadix) Build(keys [][]byte, values []interface{}) {
	t.tree = radix.New()
	for i := range keys {
		t.tree.Set(string(keys[i]), values[i])
	}
}

func (t *sauerbratenRadix) LongestPrefix(s []byte) (interface{}, bool) {
	// s is assumed contains only ASCII characters.

	var v interface{}

	tt := t.tree
	i := 0
	for i < len(s) {
		tt = tt.SubTreeWithPrefix(string([]byte{s[i]}))
		if tt == nil {
			break
		}
		v = tt.Value()
		if !strings.HasPrefix(string(s[i:]), tt.Key()) {
			break
		}
		i += len(tt.Key())
	}

	return v, i > 0
}

func (t *sauerbratenRadix) Get(s []byte) (interface{}, bool) {
	v := t.tree.Get(string(s))
	return v, v != nil
}
