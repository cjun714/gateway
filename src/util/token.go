package util

import "math/rand"

// GenToken is used to generate a 32 bytes token
func GenToken() string {
	var bts [32]byte
	for i := 0; i < 32; i++ {
		if rand.Intn(2) == 1 {
			bts[i] = byte('a') + byte(rand.Intn(25))
		} else {
			bts[i] = byte('A') + byte(rand.Intn(25))
		}
	}

	return string(bts[0:32])
}
