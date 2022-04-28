package main

import (
	"fmt"
	"unsafe"
)

type carefreeStructure struct {
	Int   bool    // 1 byte
	Float float64 // 8 bytes
	Bool  int32   // 4 bytes
}

type worriedStructure struct {
	Float float64 // 8 bytes
	Int   int32   // 4 bytes
	Bool  bool    // 1 byte
}

func main() {
	fmt.Println(unsafe.Sizeof(carefreeStructure{})) // unordered 24 bytes
	fmt.Println(unsafe.Sizeof(worriedStructure{}))  // ordered 16 bytes
}
