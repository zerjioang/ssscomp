package dghv

import (
	"fmt"
	"github.com/zerjioang/ssscomp/lib/common"
	"math"
	"math/big"
	"math/rand"
	"time"
)

// len(c) = aprox 2 * 10^7 bits
type Schema struct {
	// DGHV schema parameters
	lam   int
	rho   int64
	eta   int
	gamma int
	alpha int64
	tao   int
	rhoo  int64
	pk    privateKey
	sk    *big.Int
}

type privateKey struct {
	seed int64
	x0   *big.Int   // secret key
	di   []*big.Int // secret key pairs
}

// creates a new schema with Parameters from Public Key Compression paper
// toy_parameters = {"lam": 42, "rho": 27, "eta": 1026, "gamma": 150000, "alpha": 200, "tao": 158}
// small_parameters = {"lam": 52, "rho": 41, "eta": 1558, "gamma": 830000, "alpha": 1476, "tao": 572}
// medium_parameters = {"lam": 65, "rho": 56, "eta": 2128, "gamma": 4200000, "alpha": 2016, "tao": 1972}
func Default() Schema {
	return NewSchema(42, 27, 1026, 150000, 158, 200)
}

// creates a new schema
func NewSchema(lam, rho, eta, gamma, tao, alpha int) Schema {
	return Schema{
		lam:   lam,
		rho:   int64(rho),
		eta:   eta,
		gamma: gamma,
		tao:   tao,
		alpha: int64(alpha),
	}
}

func (s *Schema) Keygen() error {
	// set random seed based on current time
	rand.Seed(time.Now().Unix())

	// generate a random prime number p as secret key
	upperBound := bigPow(big.NewInt(2), big.NewInt(int64(s.eta)))
	lowerBound := bigPow(big.NewInt(2), big.NewInt(int64(s.eta-1)))
	fmt.Println(common.BigIntAsDecimal(upperBound))
	fmt.Println(common.BigIntAsDecimal(lowerBound))
	p, genErr := common.RandomPrimeBetween(lowerBound, upperBound)
	if genErr != nil {
		return genErr
	}
	fmt.Println(common.BigIntAsDecimal(p))
	fmt.Println("finding random odd q0")
	q0 := big.NewInt(0)
	// Procedure from KeyGen(1^lam)
	// Choose random odd q0 between 0 and max
	dividend := bigPow(big.NewInt(2), big.NewInt(int64(s.gamma)))
	q0Max := big.NewInt(0).Quo(dividend, p)
	for big.NewInt(0).Mod(q0, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		q0 = common.RandomBigNumberLessThan(q0Max)
	}

	fmt.Println("q0 found")
	fmt.Println(common.BigIntAsDecimal(q0))

	fmt.Println("computing x0...")
	x0 := big.NewInt(0).Mul(q0, p)

	// Set seed for recovery of X_i at encryption
	var seed int64 = 0
	rand.Seed(seed)

	// 0 <= X_i < 2^gamma
	fmt.Println("generating random list of values xi")
	xi := make([]*big.Int, s.tao)
	xiMax := bigPow(big.NewInt(2), big.NewInt(int64(s.gamma)))
	for i := 0; i < s.tao; i++ {
		xi[i] = common.RandomBigNumberLessThan(xiMax)
	}

	// Continue normal random
	rand.Seed(time.Now().Unix())

	// 0 <= E < 2^(lam+eta)//p
	fmt.Println("generating random list of values E")
	e := make([]*big.Int, s.tao)
	dividend2 := bigPow(big.NewInt(2), big.NewInt(int64(s.lam+s.eta)))
	maxE := big.NewInt(0).Quo(dividend2, p)
	for i := 0; i < s.tao; i++ {
		e[i] = common.RandomBigNumberLessThan(maxE)
	}

	// -2^rho < r < 2^rho
	fmt.Println("generating random list of ri")
	ri := make([]*big.Int, s.tao)
	emin := big.NewInt(int64(math.Pow(-2, float64(s.rho))) + 1)
	emax := big.NewInt(int64(math.Pow(2, float64(s.rho))))
	fmt.Println("min:", common.BigIntAsDecimal(emin))
	fmt.Println("max:", common.BigIntAsDecimal(emax))
	for i := 0; i < s.tao; i++ {
		ri[i] = common.RandomBigBetween(emin, emax)
	}
	// Construct d_i
	fmt.Println("generating list di")
	di := make([]*big.Int, s.tao)
	for i := 0; i < s.tao; i++ {
		// di[i] = xi[i]%p + e[i]*p - ri[i]
		ximodp := big.NewInt(0).Mod(xi[i], p)
		eimulp := big.NewInt(0).Mul(e[i], p)
		ximodpaddeimulp := big.NewInt(0).Add(ximodp, eimulp)
		di[i] = big.NewInt(0).Sub(ximodpaddeimulp, ri[i])
	}

	// return private and secret key pairs
	s.pk = privateKey{
		seed: seed,
		x0:   x0,
		di:   di,
	}
	s.sk = p
	return nil
}

func (s *Schema) Encrypt(m int) (*big.Int, error) {
	seed := s.pk.seed
	x0 := s.pk.x0
	di := s.pk.di

	// Recover X_i from seed
	rand.Seed(seed)
	fmt.Println("computing Xi list")
	Xi := make([]*big.Int, s.tao)
	XiMax := bigPow(big.NewInt(2), big.NewInt(int64(s.gamma)))
	for i := 0; i < s.tao; i++ {
		Xi[i] = common.RandomBigNumberLessThan(XiMax)
	}

	// #Generate x_i from X_i and d_i
	fmt.Println("generating xi from Xi and di")
	xi := make([]*big.Int, s.tao)
	for i := 0; i < s.tao; i++ {
		xi[i] = big.NewInt(0).Sub(Xi[i], di[i])
	}

	// randomize time seed again
	rand.Seed(time.Now().Unix())

	fmt.Println("computing bi list")
	bi := make([]*big.Int, s.tao)
	biMax := bigPow(big.NewInt(2), big.NewInt(s.alpha))
	for i := 0; i < s.tao; i++ {
		bi[i] = common.RandomBigNumberLessThan(biMax)
	}

	s.rhoo = s.rho + s.alpha // + w(log(lambda))
	fmt.Println("generating random r between -2^rho+1 and 2^rho")
	// r = ZZ.random_element(-2^(self.rho_) + 1, 2^(self.rho_))
	r := common.RandomBigBetween(
		big.NewInt(0).Add(
			bigPow(big.NewInt(-2), big.NewInt(s.rhoo)),
			big.NewInt(1),
		),
		bigPow(big.NewInt(2), big.NewInt(s.rhoo)))

	// initialize cipher
	c := big.NewInt(0)
	for i := 0; i < s.tao; i++ {
		// c = (c+x_i[i]*b_i[i])
		c = big.NewInt(0).Add(c, big.NewInt(0).Mul(xi[i], bi[i]))
	}
	// c = (m+2*r+2*c)%x_0
	var2r := big.NewInt(0).Mul(big.NewInt(2), r)
	var2c := big.NewInt(0).Mul(big.NewInt(2), c)
	var2rc := big.NewInt(0).Add(var2r, var2c)
	finalc := big.NewInt(0).Add(big.NewInt(int64(m)), var2rc)
	return big.NewInt(0).Mod(finalc, x0), nil
}

func (s *Schema) Decrypt(cipher *big.Int) *big.Int {
	// plaintext = (c % p) % 2
	rounded := big.NewInt(0).Div(cipher, s.sk)
	mul := big.NewInt(0).Mul(rounded, s.sk)
	cminus := big.NewInt(0).Sub(cipher, mul)
	return big.NewInt(0).Mod(cminus, big.NewInt(2))
}

func bigPow(base, exp *big.Int) *big.Int {
	return big.NewInt(0).Exp(base, exp, nil)
}
