# bitmap
Yes! It's Go/Golang's bitmap(bitset) function!

This function is achieved by the official package: "math/big" (big.Int).

It's high efficiency and easy to use.


## Install

	go get github.com/plhwin/bitmap

## Example

	package main

	import (
		"fmt"
		"github.com/plhwin/bitmap"
	)

	func main() {
		x := bitmap.New().Set(0).Set(1).Set(3).Set(5).Set(7).Set(9)
		y := bitmap.New().Set(2).Set(4).Set(5).Set(6).Set(8)
		z := bitmap.New() // z is a new empty bitmap

		// remove offset 5 from x, now x's offset is [0 1 3 7 9], bitmap is 1010001011
		x.Clear(5)

		// set new offset to y agin, now y's offset is [2 4 5 6 8 9], bitmap is 1101110100
		y.Set(9)

		//find out the same bitset from x and y
		and := x.And(y)
		// remove the same bitset from x and y
		or := x.Or(y)
		// find out the different bitset from (x and y)+(y and x)
		xor := x.Xor(y)
		// find out the different bitset from (x and y)
		andnot := x.AndNot(y)

		fmt.Println("x:", x, x.GetAllSetBits(true)) // if false, it will be return a random slice
		fmt.Println("y:", y, y.GetAllSetBits(true))
		fmt.Println("z:", z, z.GetAllSetBits(true), z.BitLen(), z.IsEmpty()) // this bitmap's len is 0, and z.IsEmpty() is true
		fmt.Println("and:", and, and.GetAllSetBits(true))
		fmt.Println("or:", or, or.GetAllSetBits(true))
		fmt.Println("xor:", xor, xor.GetAllSetBits(true))
		fmt.Println("andnot:", andnot, andnot.GetAllSetBits(true))
		
		// here if you want to check a offset in x, you can do it like this:
		fmt.Println("if offset 0 in x?", x.Test(0)) //true
		fmt.Println("if offset 5 in x?", x.Test(5)) //false, Because we run the code 'x.Clear(5)' above.
	}

code output:

	x: 1010001011 [0 1 3 7 9]
	y: 1101110100 [2 4 5 6 8 9]
	z: 0 [] 0 true
	and: 1000000000 [9]
	or: 1111111111 [0 1 2 3 4 5 6 7 8 9]
	xor: 111111111 [0 1 2 3 4 5 6 7 8]
	andnot: 10001011 [0 1 3 7]
	if offset 0 in x? true
	if offset 5 in x? false


##Tips:

let's see the output:

	x: 1010001011 [0 1 3 7 9]

please from right to left to see the bitmap `1010001011`, so you can find the corresponding relation between the offset slice `[0 1 3 7 9]` and bitmap `1010001011`,I told you to avoid your strange.


##Others
based on official package "math/big" (big.Int) implementation,the length of the bitmap that can be stored depends on the size of your memory. [more about "math/big"](https://golang.org/pkg/math/big/)

when the bitmap's length exceeds 10 billion, the performance is beginning to decline,combined with the most of business scene,I think this package can be up to more than 99% of the product scene.

if you have any problems, please submit issues here, I'll promptly pay attention and reply, thanks
