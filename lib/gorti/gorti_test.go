package gorti

import (
	"math/big"
	"testing"

	"github.com/zerjioang/ssscomp/lib/common"
)

func TestGortiScheme(t *testing.T) {
	t.Run("encrypt-decrypt", func(t *testing.T) {
		gortiSchema := NewGorti(30)
		t.Log("encrypting message: 1000")
		encryptedInt := gortiSchema.encrypt(big.NewInt(1000))
		t.Log("decrypting message: ", common.BigIntAsHex(encryptedInt))
		decryptedInt := gortiSchema.decrypt(encryptedInt)
		t.Log("decryption result: ", decryptedInt)
	})
	t.Run("encrypt-add-decrypt-same-schema", func(t *testing.T) {
		gortiSchemaA := NewGorti(30)
		encryptedA := gortiSchemaA.encrypt(big.NewInt(50))
		encryptedB := gortiSchemaA.encrypt(big.NewInt(40))
		sumab := big.NewInt(0).Add(encryptedA, encryptedB)
		decyptedContent := gortiSchemaA.decrypt(sumab)
		t.Log(decyptedContent)
	})
	t.Run("encrypt-sub-decrypt-same-schema", func(t *testing.T) {
		gortiSchemaA := NewGorti(30)
		encryptedA := gortiSchemaA.encrypt(big.NewInt(50))
		encryptedB := gortiSchemaA.encrypt(big.NewInt(40))
		sumab := big.NewInt(0).Sub(encryptedA, encryptedB)
		decyptedContent := gortiSchemaA.decrypt(sumab)
		t.Log(decyptedContent)
	})
	t.Run("encrypt-mul-decrypt-same-schema", func(t *testing.T) {
		gortiSchemaA := NewGorti(30)
		encryptedA := gortiSchemaA.encrypt(big.NewInt(50))
		encryptedB := gortiSchemaA.encrypt(big.NewInt(40))
		sumab := big.NewInt(0).Mul(encryptedA, encryptedB)
		decyptedContent := gortiSchemaA.decrypt(sumab)
		t.Log(decyptedContent)
	})
	// note: below test will never pass since two differente schemas are being used
	t.Run("encrypt-add-decrypt-different-schema", func(t *testing.T) {
		gortiSchemaA := NewGorti(30)
		gortiSchemaB := NewGorti(30)
		// schemas must share random number R
		gortiSchemaB.r = gortiSchemaA.r

		encryptedA := gortiSchemaA.encrypt(big.NewInt(50))
		encryptedB := gortiSchemaB.encrypt(big.NewInt(40))
		sumab := big.NewInt(0).Add(encryptedA, encryptedB)
		decyptedContentA := gortiSchemaA.decrypt(sumab)
		decyptedContentB := gortiSchemaB.decrypt(sumab)
		t.Log(decyptedContentA)
		t.Log(decyptedContentB)
	})
}
