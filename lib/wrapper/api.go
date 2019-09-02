//go build -buildmode=c-shared -o secure_computation_ssscomp_0.1.so api.go

//libadd.go
package main

import "C"
import (
	"fmt"
	"github.com/zerjioang/ssscomp/lib/simple"
	"unsafe"
)

const (
	v      = 0.1
	banner = `
  ______ ______ ______ ____  ____   _____ ______  
 /  ___//  ___//  ___// ___\/  _ \ /     \\____ \ 
 \___ \ \___ \ \___ \\  \__(  <_> )  Y Y  \  |_> >
/____  >____  >____  >\___  >____/|__|_|  /   __/ 
     \/     \/     \/     \/            \/|__|    
  
  Secret Sharing & Secure Computation Library

  Integrity :         None
  Version   :         0.1
`
)

//export hello
func hello() {
	fmt.Println(banner)
}

//export version
func version() float64 {
	return v
}

//export Add
func Add(a, b int) int { return a + b }

//export new_smpc_additive
func new_smpc_additive(participants C.int) uintptr {
	fmt.Println("creating new spc additive schema")
	sc, _ := simple.NewSimpleAdditiveScheme(int(participants))
	return  uintptr(unsafe.Pointer(&sc))
}

//export new_smpc_shamir
func new_smpc_shamir(participants C.int, minimum C.int) {

}

func main() {
}
