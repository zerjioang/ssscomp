package common

import (
	"fmt"
	"math/big"
)

func BigIntAsHex(v *big.Int) string {
	return fmt.Sprintf("%x", v) // or %X or upper case
}
func BigIntAsDecimal(v *big.Int) string {
	return fmt.Sprintf("%d", v) // or %X or upper case
}
