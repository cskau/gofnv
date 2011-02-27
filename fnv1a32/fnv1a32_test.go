// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fnv1a32

import (
	"io"
	"testing"
)

type _FNV1a32Test struct {
	out uint32
	in  string
}

var golden = []_FNV1a32Test{
	{0x811c9dc5, ""},
	{0xe40c292c, "a"},
	{0x4d2505ca, "ab"},
	{0x1a47e90b, "abc"},
	{0xce3479bd, "abcd"},
	{0x749bcf08, "abcde"},
	{0xff478a2a, "abcdef"},
	{0x2a9eb737, "abcdefg"},
	{0x76daaa8d, "abcdefgh"},
	{0xfe3b04ec, "abcdefghi"},
	{0xbce81ef2, "abcdefghij"},
	{0x888cf140, "Discard medicine more than two years old."},
	{0xeb0bc6ff, "He who has a shady past knows that nice guys finish last."},
	{0x0055b09d, "I wouldn't marry him with a ten foot pole."},
	{0x14c8722b, "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave"},
	{0xb195175e, "The days of the digital watch are numbered.  -Tom Stoppard"},
	{0xf9b6a42c, "Nepal premier won't resign."},
	{0xb562ff2b, "For every action there is an equal and opposite government program."},
	{0xf7a89ecc, "His money is twice tainted: 'taint yours and 'taint mine."},
	{0x3abb1b6d, "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977"},
	{0x2642aa86, "It's a tiny change to the code and not completely disgusting. - Bob Manchek"},
	{0xfefebead, "size:  a.out:  bad magic"},
	{0xbc59c00f, "The major problem is with sendmail.  -Mark Horton"},
	{0xeb610242, "Give me a rock, paper and scissors and I will move the world.  CCFestoon"},
	{0xbe861c0b, "If the enemy is within range, then so are you."},
	{0x9e6ee647, "It's well we cannot hear the screams/That we create in others' dreams."},
	{0xb19ff917, "You remind me of a TV show, but that's all right: I watch it anyway."},
	{0xf9ca0941, "C is as portable as Stonehedge!!"},
	{0x8feec41e, "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley"},
	{0xbb9a37b5, "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule"},
	{0x84271426, "How can you write a big system without C++?  -Paul Glick"},
	{0x3d767576, "'Invariant assertions' is the most elegant programming technique!  -Tom Szymanski"},
}

func TestGolden(t *testing.T) {
	for i := 0; i < len(golden); i++ {
		g := golden[i]
		c := New()
		io.WriteString(c, g.in)
		s := c.Sum32()
		if s != g.out {
			t.Errorf("fnv1a32(%s) = 0x%x want 0x%x", g.in, s, g.out)
			t.FailNow()
		}
	}
}
