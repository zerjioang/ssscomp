//go build -buildmode=c-shared -o secure_computation_s3go_0.1.so api.go

//libadd.go
package main

import "C"
import "fmt"

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

func example(path *C.char) {

}

func main() {
}
