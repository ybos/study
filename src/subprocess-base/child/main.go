package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Child process:", os.Getpid())
}
