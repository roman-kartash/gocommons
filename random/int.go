package random

import (
	"crypto/rand"
	"math/big"
)

// Int returns number in bounds [min, max).
// Panic if min > max or min < 0.
//
//nolint:predeclared
func Int(min, max int) int {
	if min > max {
		panic("min can not be more than max")
	}

	if min < 0 {
		panic("min can not be less than zero")
	}

	diff := max - min

	if diff == 0 {
		return min
	}

	bigGen, err := rand.Int(rand.Reader, big.NewInt(int64(diff)))
	if err != nil {
		panic(err)
	}

	return int(bigGen.Int64()) + min
}
