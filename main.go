package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Printf("Could not open file: %s\n", err)
	}
	t, err := decode(f)
	if err != nil {
		fmt.Printf("Could not open torrent: %s\n", err)
	}
	fmt.Printf("Torrent: %v\n", t)
}
