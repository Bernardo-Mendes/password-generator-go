package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	UPPER            = "ABCDEFGHIJKLMNOPQRSTUVWXYZÇ"
	LOWER            = "abcdefghijklmnopqrstuvwxyzç"
	NUMBER           = "0123456789"
	SPECIAL          = "!@#$%^&*()_+-=[]{}|;:,.<>?"
	MAX_LENGHT int64 = 32
)

func generateRand(cap int64) int64 {
	r, _ := rand.Int(rand.Reader, big.NewInt(cap))
	return r.Int64()
}

func numberOfChars() (int64, int64, int64, int64) {
	types := int64(4)
	min := int64(1)

	remaining := MAX_LENGHT - min*types

	var nUpper, nLower, nSpecial, nNumber int64 = 1, 1, 1, 1

	if remaining > 0 {
		add := generateRand(remaining)
		nUpper += add
		remaining -= add

		if remaining > 0 {
			add := generateRand(remaining)
			nLower += add
			remaining -= add
		}

		if remaining > 0 {
			add := generateRand(remaining)
			nSpecial += add
			remaining -= add
		}

		nNumber += remaining
	}

	return nLower, nUpper, nSpecial, nNumber
}

func shuffleBytes(b []byte) {
	// Fisher-Yates
	for i := len(b) - 1; i > 0; i-- {
		j := int(generateRand(int64(i + 1)))
		b[i], b[j] = b[j], b[i]
	}
}

func main() {
	/*
		32 characters (UPPER, lower, special, number)
		Random:
			DONE - How many characters of each type will exist in the password (at least 1 of each)
			DONE - Choose randomly from each array the chosen characters
			DONE - Mix these characters
	*/

	nLower, nUpper, nSpecial, nNumber := numberOfChars()

	arrLower := make([]byte, nLower)
	arrUpper := make([]byte, nUpper)
	arrSpecial := make([]byte, nSpecial)
	arrNumber := make([]byte, nNumber)

	for i := 0; i < int(nLower); i++ {
		randIndex := generateRand(int64(len(LOWER)))
		arrLower[i] = LOWER[randIndex]
	}

	for i := 0; i < int(nUpper); i++ {
		randIndex := generateRand(int64(len(UPPER)))
		arrUpper[i] = UPPER[randIndex]
	}

	for i := 0; i < int(nSpecial); i++ {
		randIndex := generateRand(int64(len(SPECIAL)))
		arrSpecial[i] = SPECIAL[randIndex]
	}

	for i := 0; i < int(nNumber); i++ {
		randIndex := generateRand(int64(len(NUMBER)))
		arrNumber[i] = NUMBER[randIndex]
	}

	arrFinal := append(append(append(arrLower, arrUpper...), arrSpecial...), arrNumber...)

	shuffleBytes(arrFinal)
	fmt.Print(string(arrFinal))
}
