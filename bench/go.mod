module github.com/acomagu/trie/bench

go 1.20

require (
	github.com/acomagu/trie/v2 v2.0.0
	github.com/armon/go-radix v1.0.0
	github.com/derekparker/trie v0.0.0-20190322172448-1ce4922c7ad9
	github.com/dghubble/trie v0.0.0-20190512033633-6d8e3fa705df
	github.com/hashicorp/go-immutable-radix v1.1.0
	github.com/kkdai/radix v0.0.0-20181128172204-f0c88ccaf15e
	github.com/matryer/is v1.2.0
	github.com/sauerbraten/radix v0.0.0-20150210222551-4445e9cd8982
	github.com/tchap/go-patricia v2.3.0+incompatible
)

require (
	github.com/hashicorp/golang-lru v0.5.0 // indirect
	golang.org/x/exp v0.0.0-20230310171629-522b1b587ee0 // indirect
)

replace github.com/acomagu/trie/v2 v2.0.0 => ../
