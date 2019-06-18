package bench

import (
	"github.com/tchap/go-patricia/patricia"
)

type tchapGoPatricia struct {
	tree *patricia.Trie
}

func (t *tchapGoPatricia) Name() string {
	return "github.com/tchap/go-patricia/patricia"
}

func (t *tchapGoPatricia) Build(keys [][]byte, values []interface{}) {
	t.tree = patricia.NewTrie()
	for i := range keys {
		t.tree.Insert(patricia.Prefix(keys[i]), values[i])
	}
}

func (t *tchapGoPatricia) Get(s []byte) (interface{}, bool) {
	itm := t.tree.Get(s)
	return itm, itm != nil
}

func (t *tchapGoPatricia) LongestPrefix(src []byte) (interface{}, bool) {
	var v interface{}
	for end := 1; end <= len(src); end++ {
		got := t.tree.Get(patricia.Prefix(src[:end]))
		if got != nil {
			v = got
		}
		if !t.tree.MatchSubtree(patricia.Prefix(src[:end])) {
			break
		}
	}
	return v, v != nil
}
