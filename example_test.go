package trie_test

import (
	"fmt"

	"github.com/acomagu/trie"
)

func Example_Match() {
	keys := [][]byte{
		[]byte("ab"),
		[]byte("abc"),
		[]byte("abd"),
	}
	values := []interface{}{1, 2, 3}
	t := trie.New(keys, values)

	v, ok := t.Trace([]byte("abc")).Terminal()
	fmt.Println(v, ok) // Output: 2 true
}

func Example_LongestPrefix() {
	keys := [][]byte{
		[]byte("ab"),
		[]byte("abc"),
		[]byte("abd"),
	}
	values := []interface{}{1, 2, 3}
	t := trie.New(keys, values)

	var v interface{}
	var match bool
	for _, c := range []byte("abcxxx") {
		if t = t.TraceByte(c); t == nil {
			break
		}
		if vv, ok := t.Terminal(); ok {
			v = vv
			match = true
		}
	}

	fmt.Println(v, match) // Output: 2 true
}

func Example_Terminal() {
	keys := [][]byte{ []byte("aa") }
	values := []interface{}{1}
	t := trie.New(keys, values)

	t = t.TraceByte('a') // a
	fmt.Println(t.Terminal())

	t = t.TraceByte('a') // aa
	fmt.Println(t.Terminal())

	t = t.TraceByte('a') // aaa
	fmt.Println(t.Terminal())

	// Output:
	// <nil> false
	// 1 true
	// <nil> false
}
