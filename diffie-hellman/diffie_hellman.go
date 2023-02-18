package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {

	max := new(big.Int).Sub(p, big.NewInt(2))
	randomNumber, _ := rand.Int(rand.Reader, max)

	return randomNumber.Add(randomNumber, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	A := PublicKey(a, p, g)
	return a, A
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
