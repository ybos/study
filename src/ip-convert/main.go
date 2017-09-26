package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	ip := os.Args[1]
	addr := net.ParseIP(ip)

	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The Address is", addr.String())
	}

	os.Exit(0)
}
