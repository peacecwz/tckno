package tckno

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}
