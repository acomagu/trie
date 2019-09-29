# trie: The fast and flexible Trie Tree implementation

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/acomagu/trie)

The [Trie Tree](https://wikipedia.org/wiki/Trie) implementation in Go. It has flexible interface and works fast as Radix Tree implementation.

## Benchmark

Run `make bench` to run it locally.

### Exact Match

The task is to determine whether a string matches one of all Wikipedia titles.

| Package                      |                Time |    Objects Allocated |
| ---------------------------- | -------------------:| --------------------:|
| **acomagu/trie**             |      **1090 ns/op** |      **0 allocs/op** |
| sauerbraten/radix            |          2445 ns/op |          0 allocs/op |
| dghubble/trie                |          2576 ns/op |          0 allocs/op |
| hashicorp/go-immutable-radix |          3660 ns/op |          0 allocs/op |
| derekparker/trie             |          4010 ns/op |          0 allocs/op |
| armon/go-radix               |         11745 ns/op |          0 allocs/op |
| kkdai/radix                  |         18809 ns/op |          0 allocs/op |
| tchap/go-patricia/patricia   |         21498 ns/op |          0 allocs/op |

### Longest Prefix

The task is to answer which of all Wikipedia titles can be the longest prefix of a string.

| Package                      |                Time |    Objects Allocated |
| ---------------------------- | -------------------:| --------------------:|
| **acomagu/trie**             |       **140 ns/op** |      **0 allocs/op** |
| hashicorp/go-immutable-radix |           159 ns/op |          0 allocs/op |
| tchap/go-patricia/patricia   |           252 ns/op |          0 allocs/op |
| derekparker/trie             |          2374 ns/op |          0 allocs/op |
| sauerbraten/radix            |       3264938 ns/op |          0 allocs/op |
| armon/go-radix               |      22129827 ns/op |          1 allocs/op |

(dghubble/trie and kkdai/radix don't have way to do.)

### Build

The task is to prepare Trie/Radix Tree with all of the Wikipedia titles.

| Package                      |                Time |    Objects Allocated |
| ---------------------------- | -------------------:| --------------------:|
| sauerbraten/radix            |     118959250 ns/op |     408564 allocs/op |
| **acomagu/trie**             | **542902000 ns/op** | **421906 allocs/op** |
| dghubble/trie                |     609406300 ns/op |    1136281 allocs/op |
| derekparker/trie             |    1046705400 ns/op |    1801539 allocs/op |
| armon/go-radix               |    1750312500 ns/op |    1446050 allocs/op |
| kkdai/radix                  |    2280362300 ns/op |    1742841 allocs/op |
| tchap/go-patricia/patricia   |    2898335700 ns/op |    1150947 allocs/op |
| hashicorp/go-immutable-radix |    7614342400 ns/op |   45097986 allocs/op |

## Examples

The common preparation for each examples:

```Go
keys := [][]byte{
	[]byte("ab"),
	[]byte("abc"),
	[]byte("abd"),
}
values := []interface{}{1, 2, 3}
t := trie.New(keys, values)
```

`New()` takes keys and values as the arguments. `values[i]` is the *value* of the corresponding key, `keys[i]`.

### Exact Match

```Go
v, ok := t.Trace([]byte("abc")).Terminal()
fmt.Println(v, ok) // => 2 true
```

[Playground](https://play.golang.org/p/zi6qql1x0N_y)

### Longest Prefix

```Go
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

fmt.Println(v, match) // => 2 true
```

[Playground](https://play.golang.org/p/kMfsi15FItP)

No special function to get longest prefix because it can be implemented yourself easily using the existing methods.
