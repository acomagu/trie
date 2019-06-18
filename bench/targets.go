package bench

type target interface {
	Name() string
	Build([][]byte, []interface{})
	Get([]byte) (interface{}, bool)
}

type targetHasLongestPrefix interface {
	target
	LongestPrefix([]byte) (interface{}, bool)
}

var targetFactories = []func() target{
	func() target { return new(acomaguTrie) },
	func() target { return new(armonGoRadix) },
	func() target { return new(hashicorpGoImmutableRadix) },
	func() target { return new(tchapGoPatricia) },
	func() target { return new(kkdaiRadix) },
	func() target { return new(sauerbratenRadix) },
	func() target { return new(dghubbleTrie) },
	func() target { return new(derekparkerTrie) },
}
