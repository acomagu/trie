package trie

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestTrace(t *testing.T) {
	is := is.New(t)

	type sc struct {
		path [][]byte
		nl   bool
	}
	cases := []struct {
		ss  [][]byte
		scs []sc
	}{
		{
			ss: [][]byte{
				[]byte("AM"),
				[]byte("AMD"),
				[]byte("CAD"),
				[]byte("CAM"),
				[]byte("CM"),
				[]byte("DM"),
			},
			scs: []sc{
				{[][]byte{[]byte{'A'}}, false},
				{[][]byte{[]byte{'A'}, []byte{'M'}}, false},
				{[][]byte{[]byte{'A'}, []byte{'M'}, []byte{'C'}}, true},
			},
		},
		{
			ss: [][]byte{
				[]byte("12a"),
			},
			scs: []sc{
				{[][]byte{[]byte{'1'}}, false},
				{[][]byte{[]byte{'1'}, []byte{'2'}}, false},
				{[][]byte{[]byte{'1'}, []byte{'2'}, []byte{'3'}}, true},
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			values := make([]interface{}, len(c.ss))
			for _, sc := range c.scs {
				tt := New(c.ss, values)
				for _, p := range sc.path {
					tt = tt.Trace(p)
				}
				is.Equal(tt == nil, sc.nl)
			}
		})
	}
}

func TestMatch(t *testing.T) {
	is := is.New(t)

	type entry struct {
		ss       [][]byte
		complete bool
	}
	cases := []struct {
		ss [][]byte
		es []entry
	}{
		{
			ss: [][]byte{
				[]byte("AM"),
				[]byte("AMD"),
				[]byte("CAD"),
				[]byte("CAM"),
				[]byte("CM"),
				[]byte("DM"),
			},
			es: []entry{
				{
					ss: [][]byte{
						[]byte("AM"),
					},
					complete: true,
				},
				{
					ss: [][]byte{
						[]byte{'A'},
						[]byte{'M'},
					},
					complete: true,
				},
				{
					ss: [][]byte{
						[]byte{'A'},
					},
					complete: false,
				},
			},
		},
		{
			ss: [][]byte{
				[]byte("xxx"),
				[]byte("abcd"),
				[]byte("ab"),
				[]byte("abcc"),
				[]byte("abc"),
				[]byte("xxxy"),
			},
			es: []entry{
				{
					ss:       nil,
					complete: false,
				},
				{
					ss: [][]byte{
						[]byte{'a'},
					},
					complete: false,
				},
				{
					ss: [][]byte{
						[]byte{'a'},
						[]byte{'b'},
					},
					complete: true,
				},
				{
					ss: [][]byte{
						[]byte{'a'},
						[]byte{'b'},
						[]byte{'c'},
					},
					complete: true,
				},
				{
					ss: [][]byte{
						[]byte{'a'},
						[]byte{'b'},
						[]byte{'c'},
						[]byte{'e'},
					},
					complete: false,
				},
			},
		},
		{
			ss: [][]byte{
				nil,
			},
			es: []entry{
				{
					ss: [][]byte{
						nil,
					},
					complete: true,
				},
				{
					ss: [][]byte{
						[]byte{'a'},
					},
					complete: false,
				},
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var values []interface{}
			for i := range c.ss {
				values = append(values, i)
			}
			for j, e := range c.es {
				t.Run(fmt.Sprint(j), func(t *testing.T) {
					tt := New(c.ss, values)
					for _, s := range e.ss {
						tt = tt.Trace(s)
					}
					_, ok := tt.Terminal()
					is.Equal(ok, e.complete)
				})
			}
		})
	}
}

func TestPredict(t *testing.T) {
	is := is.New(t)

	type entry struct {
		ss      [][]byte
		predict []int
	}
	cases := []struct {
		ss [][]byte
		es []entry
	}{
		{
			ss: [][]byte{
				[]byte("AA"),
				[]byte("AB"),
				[]byte("BC"),
			},
			es: []entry{
				{
					ss: [][]byte{
						[]byte("A"),
					},
					predict: []int{0, 1},
				},
				{
					ss: [][]byte{
						[]byte("AA"),
					},
					predict: []int{},
				},
				{
					ss: [][]byte{
						[]byte("B"),
					},
					predict: []int{2},
				},
			},
		},
		{
			ss: [][]byte{
				[]byte("aaa"),
			},
			es: []entry{
				{
					ss: [][]byte{
						[]byte("a"),
					},
					predict: []int{0},
				},
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var vs []interface{}
			for is := range c.ss {
				vs = append(vs, is)
			}
			for j, e := range c.es {
				t.Run(fmt.Sprint(j), func(t *testing.T) {
					tt := New(c.ss, vs)
					for _, s := range e.ss {
						tt = tt.Trace(s)
					}
					predict := tt.Predict()
					t.Log(predict, "==", e.predict)
					is.Equal(len(predict), len(e.predict))
					for i := range predict {
						is.Equal(predict[i], e.predict[i])
					}
				})
			}
		})
	}
}

func TestChildren(t *testing.T) {
	is := is.New(t)

	type entry struct {
		s        []byte
		children []byte
	}
	cases := []struct {
		ss [][]byte
		es []entry
	}{
		{
			ss: [][]byte{
				[]byte("AA"),
				[]byte("AB"),
				[]byte("ABC"),
				[]byte("BC"),
				[]byte("BDE"),
				[]byte("C"),
				[]byte("CC"),
			},
			es: []entry{
				{
					s:        []byte("A"),
					children: []byte{'A', 'B'},
				},
				{
					s:        []byte("AA"),
					children: []byte{},
				},
				{
					s:        []byte("B"),
					children: []byte{'C', 'D'},
				},
				{
					s:        []byte("C"),
					children: []byte{'C'},
				},
				{
					s:        nil,
					children: []byte{'A', 'B', 'C'},
				},
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var vs []interface{}
			for is := range c.ss {
				vs = append(vs, is)
			}
			for j, e := range c.es {
				t.Run(fmt.Sprint(j), func(t *testing.T) {
					children := New(c.ss, vs).Trace(e.s).Children()
					t.Log(string(children), "==", string(e.children))
					is.Equal(len(children), len(e.children))
					for i := range children {
						is.Equal(children[i], e.children[i])
					}
				})
			}
		})
	}
}
