package dice

import "math/rand"

func Roll() int {
	return rand.Intn(6) + 1
}
