package main

import "C"

// return Go string primitive as right format for shared object creation
func SharedString(msg string) *C.char {
	return C.CString(msg)
}
