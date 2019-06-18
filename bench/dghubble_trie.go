package bench

import "github.com/dghubble/trie"

type dghubbleTrie struct {
	tree *trie.RuneTrie
}

func (t *dghubbleTrie) Name() string {
	return "github.com/dghubble/trie"
}

func (t *dghubbleTrie) Build(keys [][]byte, values []interface{}) {
	t.tree = trie.NewRuneTrie()
	for i := range keys {
		t.tree.Put(string(keys[i]), values[i])
	}
}

func (t *dghubbleTrie) Get(s []byte) (interface{}, bool) {
	v := t.tree.Get(string(s))
	return v, v != nil
}
