package dghv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDghv(t *testing.T){
	t.Run("m-bit-0", func(t *testing.T) {
		q := 1645654 // q is a large random
		p := 15645 // p is the secret-key
		r := 12 // r is a small random
		m := 0 // m is the bit plaintext (m=0 or 1)
		if 2+r >= p {
			t.Fatal("decryption will not work due to  2*r is bigger than p")
		}

		t.Log("encrypting...")
		c := q*p + 2*r + m
		// 1645654 15645 12 0 25746256854
		t.Log(q, p, r, m, c)

		t.Log("decrypting: ")
		decryptedM := c % p % 2
		t.Log("decrypted m value: ", decryptedM)
		assert.Equal(t, decryptedM, m)
	})
	t.Run("m-bit-1", func(t *testing.T) {
		q := 1645654 // q is a large random
		p := 15645 // p is the secret-key
		r := 12 // r is a small random
		m := 1 // m is the bit plaintext (m=0 or 1)
		if 2*r >= p {
			t.Fatal("decryption will not work due to  2*r is bigger than p")
		}

		t.Log("encrypting...")
		c := q*p + 2*r + m
		// 1645654 15645 12 1 25746256855
		t.Log(q, p, r, m, c)

		t.Log("decrypting: ")
		decryptedM := c % p % 2
		t.Log("decrypted m value: ", decryptedM)
		assert.Equal(t, decryptedM, m)
	})
	t.Run("addition-a-b", func(t *testing.T) {
		q1 := 1645654 // q is a large random
		q2 := 1518484 // q is a large random
		p1 := 4345345 // p is the secret-key
		p2 := 4345345 // p is the secret-key
		r1 := 12 // r is a small random
		r2 := 36 // r is a small random
		m1 := 1 // m is the bit plaintext (m=0 or 1)
		m2 := 1 // m is the bit plaintext (m=0 or 1)
		if 2*r1 >= p1 {
			t.Fatal("decryption will not work due to 2*r1 is bigger than p")
		}
		if 2*r2 >= p2 {
			t.Fatal("decryption will not work due to 2*r2 is bigger than p")
		}

		t.Log("encrypting A...")
		c1 := q1*p1 + 2*r1 + m1
		// 1645654 15645 12 1 25746256855
		t.Log(q1, p1, r1, m1, c1)

		t.Log("encrypting B...")
		c2 := q2*p2 + 2*r2 + m2
		// 1645654 15645 12 1 25746256855
		t.Log(q2, p2, r2, m2, c2)

		t.Log("decrypting A: ")
		decryptedM := (c1 % p1) % 2
		t.Log("decrypted A value: ", decryptedM)
		assert.Equal(t, decryptedM, m1)

		t.Log("decrypting B: ")
		decrypted2 := (c2 % p2) % 2
		t.Log("decrypted A value: ", decrypted2)
		assert.Equal(t, decrypted2, m2)

		ab := c1+c2
		t.Log("A + B: ", ab)
		t.Log("decrypting A+B ...")
		t.Log(ab % p1 % 2)
		t.Log(ab % p2 % 2)
		t.Log("m1+m2 (mod 2) =", m1+m2 % 2)
		t.Log("m1 xor m2 =", m1 ^ m2)
		t.Log("c1+c2 =", ab)
	})
	t.Run("addition-a-b-c", func(t *testing.T) {
		q1 := 1645654 // q is a large random
		q2 := 1518484 // q is a large random
		q3 := 2161485 // q is a large random
		p1 := 4345345 // p is the secret-key
		p2 := 1424214 // p is the secret-key
		p3 := 5649847 // p is the secret-key
		r1 := 12 // r is a small random
		r2 := 36 // r is a small random
		r3 := 65 // r is a small random
		m1 := 1 // m is the bit plaintext (m=0 or 1)
		m2 := 1 // m is the bit plaintext (m=0 or 1)
		m3 := 1 // m is the bit plaintext (m=0 or 1)
		if 2*r1 >= p1 {
			t.Fatal("decryption will not work due to  2*r1 is bigger than p")
		}
		if 2*r2 >= p2 {
			t.Fatal("decryption will not work due to  2*r2 is bigger than p")
		}
		if 2*r3 >= p3 {
			t.Fatal("decryption will not work due to  2*r2 is bigger than p")
		}

		t.Log("encrypting A...")
		c1 := q1*p1 + 2*r1 + m1
		// 1645654 15645 12 1 25746256855
		t.Log(q1, p1, r1, m1, c1)

		t.Log("encrypting B...")
		c2 := q2*p2 + 2*r2 + m2
		// 1645654 15645 12 1 25746256855
		t.Log(q2, p2, r2, m2, c2)

		t.Log("encrypting C...")
		c3 := q3*p3 + 3*r3 + m3
		// 1645654 15645 12 1 25746256855
		t.Log(q3, p3, r3, m3, c3)

		t.Log("decrypting A: ")
		decryptedM := (c1 % p1) % 2
		t.Log("decrypted A value: ", decryptedM)
		assert.Equal(t, decryptedM, m1)

		t.Log("decrypting B: ")
		decrypted2 := (c2 % p2) % 2
		t.Log("decrypted A value: ", decrypted2)
		assert.Equal(t, decrypted2, m2)

		abc := c1+c2+c3
		t.Log("A + B + C: ", abc)
		t.Log("decrypting A+B+C ...")
		t.Log(m1+m2 % 3)
	})
}
