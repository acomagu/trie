module github.com/acomagu/trie/bench

go 1.12

require (
	github.com/acomagu/trie v0.0.0-20190618102808-a2262b8e12fe
	github.com/armon/go-radix v1.0.0
	github.com/derekparker/trie v0.0.0-20190322172448-1ce4922c7ad9
	github.com/dghubble/trie v0.0.0-20190512033633-6d8e3fa705df
	github.com/hashicorp/go-immutable-radix v1.1.0
	github.com/kkdai/radix v0.0.0-20181128172204-f0c88ccaf15e
	github.com/matryer/is v1.2.0
	github.com/sauerbraten/radix v0.0.0-20150210222551-4445e9cd8982
	github.com/tchap/go-patricia v2.3.0+incompatible
)

replace github.com/acomagu/trie v0.0.0-20190618102808-a2262b8e12fe => ../
