package main

import (
	"fmt"
	"math/rand/v2"
)

func sayHello() func() {
	threadVal := rand.IntN(10000)
	return func() {
		fmt.Println(string(threadVal))
	}
}

func main() {

}
