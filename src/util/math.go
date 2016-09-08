package util

import (
	"math/rand"
)

// return the Random number < max
func GetInt64Random(max int64) int64 {
	return rand.Int63n(max)
}

// return the Random number < max
func GetInt32Random(max int32) int32 {
	return rand.Int31n(max)
}

// return the Random number < max
func GetIntRandom(max int) int {
	return rand.Intn(max)
}
