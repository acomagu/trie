package bench

import (
	"bufio"
	"compress/gzip"
	"math/rand"
	"os"
	"sync"
	"testing"
)

// "github.com/acomagu/trie"
// "github.com/hashicorp/go-immutable-radix"
// "github.com/armon/go-radix"
// "github.com/gbrlsnchs/radix"
// kkdaiRadix "github.com/kkdai/radix"
// "github.com/tchap/go-patricia/patricia"

var readData = func() func() ([][]byte, error) {
	var once sync.Once
	var data [][]byte

	return func() ([][]byte, error) {
		var er error
		once.Do(func() {
			f, err := os.Open("enwiki-latest-all-titles-in-ns0.gz")
			if err != nil {
				er = err
				return
			}
			gr, err := gzip.NewReader(f)
			if err != nil {
				er = err
				return
			}
			scn := bufio.NewScanner(gr)
			for scn.Scan() {
				data = append(data, scn.Bytes())
			}

			// uniq
			m := make(map[string]struct{}, len(data))
			for _, d := range data {
				m[string(d)] = struct{}{}
			}
			data = data[:0]
			for d := range m {
				data = append(data, []byte(d))
			}
		})

		return data, er
	}
}()

func Benchmark(b *testing.B) {
	data, err := readData()
	if err != nil {
		panic(err)
	}

	values := make([]interface{}, 0, len(data))
	for i := range data {
		values = append(values, i)
	}

	ts := make([]target, len(targetFactories))
	b.Run("Build", func(b *testing.B) {
		for it, f := range targetFactories {
			b.Run(f().Name(), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ts[it] = f()
					ts[it].Build(data, values)
				}
			})
		}
	})

	b.Run("LongestPrefix", func(b *testing.B) {
		var l int
		for _, d := range data {
			l += len(d)
		}

		src := make([]byte, 0, l)
		for _, d := range data {
			src = append(src, d...)
		}
		rand.Shuffle(len(src), func(i, j int) {
			src[i], src[j] = src[j], src[i]
		})

		for _, t := range ts {
			if t, ok := t.(targetHasLongestPrefix); ok {
				b.Run(t.Name(), func(b *testing.B) {
					begin := 0
					for i := 0; i < b.N; i++ {
						begin %= len(src)
						_, _ = t.LongestPrefix(src[begin:])
						begin++
					}
				})
			}
		}
	})

	b.Run("Match", func(b *testing.B) {
		srcs := make([]int, 0, 10000)
		for i := 0; i < 10000; i++ {
			srcs = append(srcs, rand.Intn(len(data)))
		}

		for _, t := range ts {
			b.Run(t.Name(), func(b *testing.B) {
				si := 0
				for i := 0; i < b.N; i++ {
					si %= len(srcs)
					_, _ = t.Get(data[srcs[si]])
					si++
				}
			})
		}
	})

	// kkdai/radix: Impossible

	// q-radix: Impossible
}
