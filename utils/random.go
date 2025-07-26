package utils

import (
	"math/rand/v2"
)

func CreateRNG(seed uint64) *rand.Rand {
	return rand.New(rand.NewPCG(seed, uint64(seed/2+1)))
}

func RollDice(sides int, rng *rand.Rand) int {
	if sides < 1 {
		return 0
	}
	return rng.IntN(sides) + 1
}

func RandomAsciiChar(rng *rand.Rand) byte {
	return byte(rng.IntN(93) + 32) // 'a' to 'z'
}
func RandomBinaryChar(rng *rand.Rand) string {
	res := byte(rng.IntN(2))
	if res == 0 {
		return " "
	} else {
		return "█"
	}
}
