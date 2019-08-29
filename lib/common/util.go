package common

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	mrand "math/rand"
	"strings"
	"time"
)

const (
	defaultPrimeNumber = "13407807929942597099574024998205846127479365820592393377723561443721764030073546976801874298166903427690031858186486050853753882811946569946433649006083221"
	zeroAscii          = 48
)

var (
	prime *big.Int
)

func init() {
	//generate default prime number
	prime, _ = big.NewInt(0).SetString(defaultPrimeNumber, 10)
}

func GetDefaultPrimeNumber() *big.Int {
	return prime
}

// generates a random int in specified range
func RandomInRange(min, max int) int {
	mrand.Seed(time.Now().UnixNano())
	return mrand.Intn(max-min) + min
}

// generates a random int
func RandomInt() int {
	mrand.Seed(time.Now().UnixNano())
	return mrand.Int()
}

// generates a random set of integer values
func RandomSet(size int) []int {
	var result []int
	result = make([]int, size)
	for i := 0; i < len(result); i++ {
		result[i] = int(RandomInRange(0, 100))
	}
	return result
}

func Polynomial(x int, v []int) int {
	var result int
	for i := 0; i < len(v); i++ {
		result += v[i] * int(math.Pow(float64(x), float64(i+1)))
	}
	return result
}

/**
 * Returns a RandomPrime number from the range (0, prime-1) inclusive
**/
func RandomPrime() *big.Int {
	prime := GetDefaultPrimeNumber()
	result := big.NewInt(0).Set(prime)
	result = result.Sub(result, big.NewInt(1))
	result, _ = rand.Int(rand.Reader, result)
	return result
}

/**
 * Converts a byte array into an a 256-bit big.Int, arraied based upon size of
 * the input byte; all values are right-padded to length 256, even if the most
 * significant bit is zero.
**/
func SplitByteToInt(secret []byte) []*big.Int {
	hexData := hex.EncodeToString(secret)
	count := int(math.Ceil(float64(len(hexData)) / 64.0))

	result := make([]*big.Int, count)

	for i := 0; i < count; i++ {
		if (i+1)*64 < len(hexData) {
			result[i], _ = big.NewInt(0).SetString(hexData[i*64:(i+1)*64], 16)
		} else {
			data := strings.Join([]string{hexData[i*64:], strings.Repeat("0", 64-(len(hexData)-i*64))}, "")
			result[i], _ = big.NewInt(0).SetString(data, 16)
		}
	}

	return result
}

/**
 * Converts an array of big.Ints to the original byte array, removing any
 * least significant nulls
**/
func MergeIntToByte(secret []*big.Int) []byte {
	var hexData = ""
	for i := range secret {
		tmp := fmt.Sprintf("%x", secret[i])
		hexData += strings.Join([]string{strings.Repeat("0", 64-len(tmp)), tmp}, "")
	}

	result, _ := hex.DecodeString(hexData)
	result = bytes.TrimRight(result, "\x00")

	return result
}

/**
 * Evauluates a polynomial with coefficients specified in reverse order:
 * EvaluatePolynomial([a, b, c, d], x):
 * 		returns a + bx + cx^2 + dx^3
**/
func EvaluatePolynomial(polynomial []*big.Int, value *big.Int) *big.Int {
	prime := GetDefaultPrimeNumber()
	last := len(polynomial) - 1
	var result = big.NewInt(0).Set(polynomial[last])

	for s := last - 1; s >= 0; s-- {
		result = result.Mul(result, value)
		result = result.Add(result, polynomial[s])
		result = result.Mod(result, prime)
	}

	return result
}

/**
 * InNumbers(array, value) returns boolean whether or not value is in array
**/
func InNumbers(numbers []*big.Int, value *big.Int) bool {
	for n := range numbers {
		if numbers[n].Cmp(value) == 0 {
			return true
		}
	}
	return false
}

/**
 * Returns the big.Int number base10 in base64 representation; note: this is
 * not a string representation; the base64 output is exactly 256 bits long
**/
func ToBase64(number *big.Int) string {
	hexdata := fmt.Sprintf("%x", number)
	for i := 0; len(hexdata) < 64; i++ {
		hexdata = "0" + hexdata
	}
	bytedata, success := hex.DecodeString(hexdata)
	if success != nil {
		fmt.Println("Error!")
		fmt.Println("hexdata: ", hexdata)
		fmt.Println("bytedata: ", bytedata)
		fmt.Println(success)
	}
	return base64.URLEncoding.EncodeToString(bytedata)
}

/**
 * Returns the number base64 in base 10 big.Int representation; note: this is
 * not coming from a string representation; the base64 input is exactly 256
 * bits long, and the output is an arbitrary size base 10 integer.
 *
 * Returns -1 on failure
**/
func FromBase64(number string) *big.Int {
	bytedata, err := base64.URLEncoding.DecodeString(number)
	if err != nil {
		return big.NewInt(-1)
	}

	hexdata := hex.EncodeToString(bytedata)
	result, ok := big.NewInt(0).SetString(hexdata, 16)
	if ok == false {
		return big.NewInt(-1)
	}

	return result
}

/**
 * Computes the multiplicative inverse of the number on the field prime; more
 * specifically, number * inverse == 1; Note: number should never be zero
**/
func ModInverse(number *big.Int) *big.Int {
	prime := GetDefaultPrimeNumber()
	c := big.NewInt(0).Set(number)
	c = c.Mod(c, prime)
	pcopy := big.NewInt(0).Set(prime)
	x := big.NewInt(0)
	y := big.NewInt(0)

	c.GCD(x, y, pcopy, c)

	result := big.NewInt(0).Set(prime)

	result = result.Add(result, y)
	result = result.Mod(result, prime)
	return result
}

// Calculates the modulus of given two numbers
func Mod(x, y int) int {
	if x < 0 {
		return (x % y) + y
	}
	return x % y
}

// Calculates whether given number is prime or not
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// Calculates whether given big number string decimal representation is prime or not
func IsPrimeBig(primeStr string) bool {
	lastDigit := primeStr[len(primeStr)-1]
	if lastDigit == zeroAscii || lastDigit == zeroAscii+2 || lastDigit == zeroAscii+5 {
		return false
	}
	z := new(big.Int)
	z.SetString(primeStr, 10)
	return z.ProbablyPrime(20)
}

// Computes the GCD of two big ints
func GCD(a, b *big.Int) big.Int {
	t := big.NewInt(0)
	zero := big.NewInt(0)
	for b.Cmp(zero) != 0 {
		t.Rem(a, b)
		a, b, t = b, t, a
	}
	return *a
}
