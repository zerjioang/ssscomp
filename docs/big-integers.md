# Big Integers in Go

Package big implements arbitrary-precision arithmetic (big numbers). The following numeric types are supported: 

```txt
Int    signed integers
Rat    rational numbers
Float  floating-point numbers
```

The zero value for an Int, Rat, or Float correspond to 0. Thus, new values can be declared in the usual ways and denote 0 without further initialization.

## Big Int Problem

In Go, Big Integers can have an 'unlimited' size and are not couple with size of the uint64 in specific architecture nor any similar. However, they are limited by the memory used by the computations made with them.

```go
package main

import (
    "fmt"
    "math/big"
)

func main() {
    verybig := big.NewInt(1)
    ten := big.NewInt(10)
    for i:=0; i<1000000; i++ {
       temp := new(big.Int)
       temp.Mul(verybig, ten)
       verybig = temp
    }
    fmt.Println(verybig)
}
```

Note that iterations number set to `1000000` can make your physical memory overflow and program fail.