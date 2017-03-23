package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage %s n1 n2 n3 ...", os.Args[0])
	}
	nums := os.Args[1:]
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		fmt.Printf("%x", n)
	}
	fmt.Printf("\n")
}
