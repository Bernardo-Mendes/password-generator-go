package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	UPPER      = "ABCDEFGHIJKLMNOPQRSTUVWXYZÇ"
	LOWER      = "abcdefghijklmnopqrstuvwxyzç"
	NUMBER     = "0123456789"
	SPECIAL    = "!@#$%&"
	MAX_LENGHT = 32
)

func generate_rand(cap int64) int64 {
	r, _ := rand.Int(rand.Reader, big.NewInt(cap))
	return r.Int64()
}

func main() {
	/*
		32 characters (UPPER, lower, special, number)
		Random:
			- How many characters of each type will exist in the password (at least 1 of each)
			- Choose randomly from each array the chosen characters
			- Mix these characters
	*/
	var remaining_chars int64 = MAX_LENGHT
	var n_upper_chars int64
	var n_lower_chars int64
	var n_special_chars int64
	var n_number_chars int64

	n_upper_chars = generate_rand(remaining_chars)
	remaining_chars -= n_upper_chars
	n_lower_chars = generate_rand(remaining_chars)
	remaining_chars -= n_lower_chars
	n_special_chars = generate_rand(remaining_chars)
	n_number_chars = remaining_chars
	if remaining_chars == 0 {
		n_special_chars--
		n_number_chars = remaining_chars + 1
	}

	fmt.Printf("Upper: %d, Lower: %d, Special: %d, Number: %d\n",
		n_upper_chars, n_lower_chars, n_special_chars, n_number_chars)
}
