// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fnv1a64

import (
	"io"
	"testing"
)

type _FNV1a64Test struct {
	out uint64
	in  string
}

var golden = []_FNV1a64Test{
	{0xcbf29ce484222325, ""},
	{0xaf63dc4c8601ec8c, "a"},
	{0x089c4407b545986a, "ab"},
	{0xe71fa2190541574b, "abc"},
	{0xfc179f83ee0724dd, "abcd"},
	{0x6348c52d762364a8, "abcde"},
	{0xd80bda3fbe244a0a, "abcdef"},
	{0x406e475017aa7737, "abcdefg"},
	{0x25da8c1836a8d66d, "abcdefgh"},
	{0xfb321124e0e3a8cc, "abcdefghi"},
	{0xb9bbc7aa22d79212, "abcdefghij"},
	{0x03c96a886c594e20, "Discard medicine more than two years old."},
	{0xfa2e4b8806972f1f, "He who has a shady past knows that nice guys finish last."},
	{0x51d60253e844a45d, "I wouldn't marry him with a ten foot pole."},
	{0x6ec35c72c38a5fcb, "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave"},
	{0xed6d1b0a4d3eb67e, "The days of the digital watch are numbered.  -Tom Stoppard"},
	{0xeb0df5b3742a17cc, "Nepal premier won't resign."},
	{0xd79dce3b506735eb, "For every action there is an equal and opposite government program."},
	{0x1523532136e9906c, "His money is twice tainted: 'taint yours and 'taint mine."},
	{0xd950d12cf0f5440d, "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977"},
	{0xf7d9ade991506a06, "It's a tiny change to the code and not completely disgusting. - Bob Manchek"},
	{0x7af3dab54997816d, "size:  a.out:  bad magic"},
	{0xfc1c195545bea7ef, "The major problem is with sendmail.  -Mark Horton"},
	{0x4d7b224a96a3d382, "Give me a rock, paper and scissors and I will move the world.  CCFestoon"},
	{0x279d9ba44494814b, "If the enemy is within range, then so are you."},
	{0xc9b4e336376f3ce7, "It's well we cannot hear the screams/That we create in others' dreams."},
	{0x50a13def48f21c97, "You remind me of a TV show, but that's all right: I watch it anyway."},
	{0x7239e5776dd303c1, "C is as portable as Stonehedge!!"},
	{0x3bed51a6482fd3be, "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley"},
	{0x7258208b3f2d3df5, "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule"},
	{0x988bf2e6a72399e6, "How can you write a big system without C++?  -Paul Glick"},
	{0xf7cd2216cc072e56, "'Invariant assertions' is the most elegant programming technique!  -Tom Szymanski"},
}

func TestGolden(t *testing.T) {
	for i := 0; i < len(golden); i++ {
		g := golden[i]
		c := New()
		io.WriteString(c, g.in)
		s := c.Sum64()
		if s != g.out {
			t.Errorf("fnv1a64(%s) = 0x%x want 0x%x", g.in, s, g.out)
			t.FailNow()
		}
	}
}
