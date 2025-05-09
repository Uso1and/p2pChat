package crypto

import (
	"crypto/rand"

	"golang.org/x/crypto/curve25519"
)

var basePoint = [32]byte{9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func GenerateKeyPair() (privateKey [32]byte, publicKey [32]byte, err error) {
	_, err = rand.Read(privateKey[:])
	if err != nil {
		return
	}

	privateKey[0] &= 248
	privateKey[31] &= 127
	privateKey[31] |= 64

	curve25519.ScalarMult(&publicKey, &privateKey, &basePoint)
	return
}

func ComputeSharedSecret(privateKey [32]byte, publicKey [32]byte) [32]byte {
	var sharedSecret [32]byte
	curve25519.ScalarMult(&sharedSecret, &privateKey, &publicKey)
	return sharedSecret
}
