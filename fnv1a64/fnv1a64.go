// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package implements the FNV-1a 64-bit checksum.
// See: http://en.wikipedia.org/wiki/Fowler_Noll_Vo_hash
//  hash = FNV_offset_basis
//  for each octet_of_data to be hashed
//    hash = hash XOR octet_of_data
//    hash = hash  FNV_prime
//  return hash

package fnv1a64

import (
	"hash"
	"os"
)

const (
	fnv_offset_basis = 14695981039346656037
	fnv_prime = 1099511628211
)

const Size = 8

// digest represents the partial evaluation of a hash.
type digest struct {
	hash uint64
}

func (d *digest) Reset() { d.hash = fnv_offset_basis }

// New returns a new hash.Hash64 computing the FNV-1a-64 hash.
func New() hash.Hash64 {
	d := new(digest)
	d.Reset()
	return d
}

func (d *digest) Size() int { return Size }

// Add p to the running hash.
func update(h uint64, p []byte) uint64 {
	for i := 0; i < len(p); i++ {
		h ^= uint64(p[i])
		h *= fnv_prime
	}
	return h
}

func (d *digest) Write(p []byte) (nn int, err os.Error) {
	d.hash = update(d.hash, p)
	return len(p), nil
}

func (d *digest) Sum64() uint64 {
	return d.hash
}

func (d *digest) Sum() []byte {
	p := make([]byte, 8)
	s := d.Sum64()
	p[0] = byte(s >> 56)
	p[1] = byte(s >> 48)
	p[2] = byte(s >> 40)
	p[3] = byte(s >> 32)
	p[4] = byte(s >> 24)
	p[5] = byte(s >> 16)
	p[6] = byte(s >> 8)
	p[7] = byte(s)
	return p
}
