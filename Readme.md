# gofnv

gofnv is an implementation of the Fowler-Noll-Vo hash function in the Go programming language.

## Installation

You can install the packages by using an appropriate variant of `goinstall github.com/cskau/gofnv/fnv1a64`.
If you do this, the import statement in your go programs will be something like `import github.com/cskau/gofnv/fnv1a64`.

## Example usage

    package main
    
    import (
        "io"
        "fmt"
        "github.com/cskau/gofnv/fnv1a64"
    )
    
    func main() {
        str := "Fowler-Noll-Vo"
        hasher := fnv1a64.New()
        io.WriteString(hasher, str)
        hash := hasher.Sum64()
        fmt.Printf( "fnv1a64(\"%s\") = 0x%x\n", str, hash )
    }

To run the application, put the code in a file called hello-fnv.go and run:

    8g hello-fnv.go && 8l -o hello-fnv hello-fnv.8 && ./hello-fnv

## Further reading

For further reading about the Fowler-Noll-Vo hash function see the [wikipedia article](http://en.wikipedia.org/wiki/Fowler_Noll_Vo_hash).
Or visit Landon Noll of Fowler-Noll-Vo's [website](http://www.isthe.com/chongo/tech/comp/fnv/index.html) about the function.
