package main

import (
	"log"
)

func main() {
	var x []rune
	x = []rune("4242424242424242")
	for _, a := range x {
		log.Println(a)
	}
}
