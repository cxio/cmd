package main

import (
	"fmt"
)

// 一年的区块数量（恒星年）。
const blocks = 87661

func decrementLoop(b, bv int) {
	y := 1
	sum := 0

	for b > 0 {
		sum += b * blocks
		fmt.Printf("%d\t%-8d\t%d\n", y, sum, b)
		b = b * bv / 100
		y++
	}
}

func main() {
	fmt.Printf("年次\t累计\t\t币/块\n")
	fmt.Println("--------------------------------- (90%)")

	// 逐年递减10%
	decrementLoop(50, 90)
}
