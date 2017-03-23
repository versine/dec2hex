package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	c := flag.Bool("c", false, "Concatenate results together")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Printf("Usage %s n1 n2 n3 ...", os.Args[0])
	}

	nums := flag.Args()
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		fmt.Printf("%02x", n)
		if !*c {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}
