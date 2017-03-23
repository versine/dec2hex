package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := flag.Bool("s", false, "Whether to insert spaces between individual numbers")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Printf("Usage %s n1 n2 n3 ...", os.Args[0])
	}

	nums := flag.Args()
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		fmt.Printf("%02x", n)
		if *s {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}
