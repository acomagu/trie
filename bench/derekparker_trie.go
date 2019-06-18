package bench

import (
	"github.com/derekparker/trie"
)

type derekparkerTrie struct {
	tree *trie.Trie
}

func (t *derekparkerTrie) Name() string {
	return "github.com/derekparker/trie"
}

func (t *derekparkerTrie) Build(keys [][]byte, values []interface{}) {
	t.tree = trie.New()
	for i := range keys {
		t.tree.Add(string(keys[i]), values[i])
	}
}

func (t *derekparkerTrie) Get(s []byte) (interface{}, bool) {
	n, ok := t.tree.Find(string(s))
	if !ok {
		return nil, false
	}
	return n.Meta(), true
}

func (t *derekparkerTrie) LongestPrefix(s []byte) (interface{}, bool) {
	// s is assumed contains only ASCII characters.

	var v interface{}
	var match bool
	for end := 1; end <= len(s); end++ {
		if n, ok := t.tree.Find(string(s[:end])); ok {
			match = true
			v = n.Meta()
		}
		if !t.tree.HasKeysWithPrefix(string(s[:end])) {
			break
		}
	}

	return v, match
}
