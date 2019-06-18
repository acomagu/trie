package bench

import (
	"testing"

	"github.com/matryer/is"
)

func TestTarget_Get(t *testing.T) {
	suite := struct {
		keys  [][]byte
		cases []struct {
			s  []byte
			v  int
			ok bool
		}
	}{
		keys: [][]byte{
			[]byte("AM"),
			[]byte("AMD"),
			[]byte("CAD"),
			[]byte("CAM"),
			[]byte("CM"),
			[]byte("DM"),
		},
		cases: []struct {
			s  []byte
			v  int
			ok bool
		}{
			{
				s:  []byte("CAD"),
				v:  2,
				ok: true,
			},
			{
				s:  []byte("AMM"),
				ok: false,
			},
			{
				s:  []byte("D"),
				ok: false,
			},
		},
	}

	for _, f := range targetFactories {
		tg := f()
		t.Run(tg.Name(), func(t *testing.T) {
			var values []interface{}
			for i := range suite.keys {
				values = append(values, i)
			}
			tg.Build(suite.keys, values)

			for _, c := range suite.cases {
				t.Run(string(c.s), func(t *testing.T) {
					is := is.New(t)

					v, ok := tg.Get(c.s)
					is.Equal(ok, c.ok)
					if c.ok {
						is.Equal(v, c.v)
					}
				})
			}
		})
	}
}

func TestTarget_LongestPrefix(t *testing.T) {
	suite := struct {
		keys  [][]byte
		cases []struct {
			s  []byte
			v  int
			ok bool
		}
	}{
		keys: [][]byte{
			[]byte("AM"),
			[]byte("AMD"),
			[]byte("CAD"),
			[]byte("CAM"),
			[]byte("CM"),
			[]byte("DM"),
		},
		cases: []struct {
			s  []byte
			v  int
			ok bool
		}{
			{
				s:  []byte("CADEER"),
				v:  2,
				ok: true,
			},
			{
				s:  []byte("AMO"),
				v:  0,
				ok: true,
			},
			{
				s:  []byte("CAD"),
				v:  2,
				ok: true,
			},
			{
				s:  []byte("D"),
				ok: false,
			},
			{
				s:  []byte("AMDU"),
				v: 1,
				ok: true,
			},
		},
	}

	for _, f := range targetFactories {
		tg, ok := f().(targetHasLongestPrefix)
		if !ok {
			continue
		}

		t.Run(tg.Name(), func(t *testing.T) {
			var values []interface{}
			for i := range suite.keys {
				values = append(values, i)
			}
			tg.Build(suite.keys, values)

			for _, c := range suite.cases {
				t.Run(string(c.s), func(t *testing.T) {
					is := is.New(t)

					v, ok := tg.LongestPrefix(c.s)
					is.Equal(ok, c.ok)
					if c.ok {
						is.Equal(v, c.v)
					}
				})
			}
		})
	}
}
