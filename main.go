package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZÇ"
const LOWER = "abcdefghijklmnopqrstuvwxyzç"
const NUMBER = "0123456789"
const SPECIAL = "!@#$%&*?<>"

func generate_rand() int64 {
	r, _ := rand.Int(rand.Reader, big.NewInt(100))
	return r.Int64()
}

func main() {
	/*
		32 characters (UPPER, lower, special, number)
		Random:
			- How many characters of each type will exist in the password (at least 1 of each)
			- Choose freely from each array the chosen characters
			- Mix these characters
	*/

	random1 := generate_rand()
	fmt.Println(random1)
}
