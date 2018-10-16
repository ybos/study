package main

/**
Package io provides basic interfaces to I/O primitives. Its primary job is to wrap existing
implementations of such primitives, such as those in package os, into shared public interfaces
that abstract the functionality, plus some other related primitives.

Because these interfaces and primitives wrap lower-level operations with various implementations,
unless otherwise informed clients should not assume they are safe for parallel execution.
 */

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("This is a test message.")

	n, err := io.Copy(os.Stdout, r)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nsuccess copy ", n, "Bytes")
}
