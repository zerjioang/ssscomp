package finite

// fixed a finite field to which all secrets
// and shares belong
// In mathematics, a finite field or Galois field
// (so-named in honor of Évariste Galois) is a field
// that contains a finite number of elements. As with
// any field, a finite field is a set on which the
// operations of multiplication, addition, subtraction
// and division are defined and satisfy certain basic
// rules. The most common examples of finite fields are
// given by the integers mod p when p is a prime number.
type FiniteField uint8

// https://web.eecs.utk.edu/~jplank/plank/papers/CS-07-593/primitive-polynomial-table.txt

const (
	IntegerField FiniteField = iota
	NonPrimeField
	PrimeNumberField
	GF2
	GF4
	GF8
	GF27
	GF16
	GF64
)
