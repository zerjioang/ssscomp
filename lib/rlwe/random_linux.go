// +build linux !freebsd !cgo

/*

Instruction	 Source                                                    NIST Compliance
RDRAND	     Cryptographically secure pseudorandom number generator	   SP 800-90A
RDSEED	     Non-deterministic random bit generator	                   SP 800-90B & C (drafts)

* RDSEED: gives true random values but at slow data rate
* RDRAND: gives very good random values, good enough for all but the most demanding (e.g. tightest security) applications,
          at data rate much higher than RSEED

this file implements the RNG,
please don't use OS  based RNG  and don't use RDSEED  , processor based RNG generator
The current function is just a placeholder and must be replaced before this code goes into production
TAGS are used so as this file is not used, and should be replaced
Any attempt to cross-compile would cause this file to be skipped and thus the RND gen fixed
*/
package rlwe

import (
	"crypto/rand"
	"time"
)

var x uint64 /* The state must be seeded with a nonzero value. */

func init() {
	x = uint64(time.Now().UTC().UnixNano()) // asumming we only 20 bits of entry, since time can be guessed
	z := x % 1024
	z += 41

	buf := make([]byte, z)
	_, err := rand.Read(buf)
	if err != nil {
		panic("error: initialing RNG generator")
		return
	}

	for i := uint64(0); i < z; i++ {
		x ^= x >> buf[i] & 63 // a
		x ^= x << 25          // b
		x ^= x >> buf[i] & 63 // c

		xorshift64star()
	}

	// now lets mix in os rnd generator for time being

}

func xorshift64star() uint64 {
	x ^= x >> 12 // a
	x ^= x << 25 // b
	x ^= x >> 27 // c
	return x * uint64(2685821657736338717)
}

/* this function generates a uint64 RNG */
func RandomSafeUint64() uint64 {
	return xorshift64star()
}
