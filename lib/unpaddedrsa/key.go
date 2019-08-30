package unpaddedrsa

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

// GenerateKeyPair generates a new key pair
func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Error(err)
	}
	return privkey, &privkey.PublicKey
}
