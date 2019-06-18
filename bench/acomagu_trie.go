package bench

import (
	"github.com/acomagu/trie"
)

type acomaguTrie struct {
	tree trie.Tree
}

func (t *acomaguTrie) Name() string {
	return "github.com/acomagu/trie"
}

func (t *acomaguTrie) Build(keys [][]byte, values []interface{}) {
	t.tree = trie.New(keys, values)
}

func (t *acomaguTrie) Get(s []byte) (interface{}, bool) {
	return t.tree.Trace(s).Terminal()
}

func (t *acomaguTrie) LongestPrefix(s []byte) (interface{}, bool) {
	var v interface{}
	var match bool

	tt := t.tree
	for _, c := range s {
		tt = tt.TraceByte(c)
		if tt == nil {
			break
		}
		if vv, ok := tt.Terminal(); ok {
			v = vv
			match = true
		}
	}

	return v, match
}
