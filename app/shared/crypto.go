package shared

import "crypto/rand"

// Returns securely generated random sequence of bytes.
// It will panic if an error is encontered as the user should not continue.
// Adapted from: http://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/
func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		panic(err)
	}

	return b
}
