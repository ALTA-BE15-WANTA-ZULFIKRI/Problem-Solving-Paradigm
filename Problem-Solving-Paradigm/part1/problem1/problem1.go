package main

import "fmt"

func main() {
	A, B, C := 6, 6, 14

	for x := 1; x < A; x++ {
		for y := 1; y < A; y++ {
			for z := 1; z < A; z++ {
				if x+y+z == A && x*y*z == B && x*x+y*y+z*z == C {
					fmt.Println(x, y, z)
					return
				}
			}
		}
	}
	fmt.Println("no solution")
}
