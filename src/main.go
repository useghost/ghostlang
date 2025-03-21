package main

import (
	"fmt"
	"ghostlang/encoding"
)

func main() {
	bytes := encoding.WriteI32LE(124124999);
	fmt.Println(bytes)
	fmt.Println(encoding.ReadI32LE(bytes))

	bytes16 := encoding.WriteI16LE(500);
	fmt.Println(bytes16)
	fmt.Println(encoding.ReadI16LE(bytes16))

	bytes8 := encoding.WriteI8LE(200);
	fmt.Println(bytes8)
	fmt.Println(encoding.ReadI8LE(bytes8))

	bytes64 := encoding.WriteI64LE(1844674407370955161);
	fmt.Println(bytes64)
	fmt.Println(encoding.ReadI64LE(bytes64))

	
	// fmt.Println(encoding.ReadI16LE(bytes))
	// fmt.Println(encoding.WriteI8LE(255))
}