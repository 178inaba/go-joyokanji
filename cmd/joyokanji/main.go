package main

import (
	"fmt"
	"log"

	"github.com/178inaba/go-joyokanji"
)

func main() {
	ks, err := joyokanji.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, k := range ks {
		fmt.Printf("%#U\n", k)
	}

	fmt.Printf("Total: %d.\n", len(ks))
}
