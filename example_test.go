package trie_test

import (
	"fmt"

	"github.com/acomagu/trie/v2"
)

func Example_match() {
	keys := [][]byte{
		[]byte("ab"),
		[]byte("abc"),
		[]byte("abd"),
	}
	values := []int{1, 2, 3} // The type of value doesn't have to be int. Can be anything.
	t := trie.New(keys, values)

	v, ok := t.Trace([]byte("abc")).Terminal()
	fmt.Println(v, ok) // Output: 2 true
}

func Example_longestPrefix() {
	keys := [][]byte{
		[]byte("ab"),
		[]byte("abc"),
		[]byte("abd"),
	}
	values := []int{1, 2, 3} // The type of value doesn't have to be int. Can be anything.
	t := trie.New(keys, values)

	var v interface{}
	var match bool
	for _, c := range []byte("abcxxx") {
		if t = t.TraceOne(c); t == nil {
			break
		}
		if vv, ok := t.Terminal(); ok {
			v = vv
			match = true
		}
	}

	fmt.Println(v, match) // Output: 2 true
}

func ExampleTree_Terminal() {
	keys := [][]byte{[]byte("aa")}
	values := []int{1} // The type of value doesn't have to be int. Can be anything.
	t := trie.New(keys, values)

	t = t.TraceOne('a') // a
	fmt.Println(t.Terminal())

	t = t.TraceOne('a') // aa
	fmt.Println(t.Terminal())

	t = t.TraceOne('a') // aaa
	fmt.Println(t.Terminal())

	// Output:
	// 0 false
	// 1 true
	// 0 false
}
