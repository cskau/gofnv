// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package implements the FNV-1a 32-bit checksum.
// See: http://en.wikipedia.org/wiki/Fowler_Noll_Vo_hash
//  hash = FNV_offset_basis
//  for each octet_of_data to be hashed
//    hash = hash XOR octet_of_data
//    hash = hash  FNV_prime
//  return hash

package fnv1a32

import (
	"hash"
	"os"
)

const (
	fnv_offset_basis = 2166136261
	fnv_prime = 16777619
)

const Size = 4

// digest represents the partial evaluation of a hash.
type digest struct {
	hash uint32
}

func (d *digest) Reset() { d.hash = fnv_offset_basis }

// New returns a new hash.Hash32 computing the FNV-1a-64 hash.
func New() hash.Hash32 {
	d := new(digest)
	d.Reset()
	return d
}

func (d *digest) Size() int { return Size }

// Add p to the running hash.
func update(h uint32, p []byte) uint32 {
	for i := 0; i < len(p); i++ {
		h ^= uint32(p[i])
		h *= fnv_prime
	}
	return h
}

func (d *digest) Write(p []byte) (nn int, err os.Error) {
	d.hash = update(d.hash, p)
	return len(p), nil
}

func (d *digest) Sum32() uint32 {
	return d.hash
}

func (d *digest) Sum() []byte {
	p := make([]byte, 4)
	s := d.Sum32()
	p[0] = byte(s >> 24)
	p[1] = byte(s >> 16)
	p[2] = byte(s >> 8)
	p[3] = byte(s)
	return p
}
