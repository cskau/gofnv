package main

import (
	"github.com/cskau/gofnv/fnv1a64"
)

func main() {
	str := "Fowler-Noll-Vo"
	hasher := fnv1a64.New()
	io.WriteString(hasher, str)
	hash := hasher.Sum64()
	println( "fnv1a64(%s) = 0x%x", str, hash )
}
