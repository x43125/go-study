package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 20; i++ {
		n := rand.Intn(15) + 1
		fmt.Println(n)
	}
}
