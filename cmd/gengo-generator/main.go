package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/178inaba/go-joyokanji"
)

func main() {
	ks, err := joyokanji.List()
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Printf("%c%c\n", ks[rand.Intn(len(ks))], ks[rand.Intn(len(ks))])
}
