package main

import (
	"fmt"
)

func generateTestArray(n, m int) {

	iterator := 0
	for i := 0; i < n; i++ {
		fmt.Print("{")
		for j := 0; j < m; j++ {
			if j == 0 {
				fmt.Print(iterator)
			} else {
				fmt.Print(",", iterator)

			}
			iterator++
		}
		fmt.Println("},")
	}
	fmt.Println()
}

func generateTestArrayWithOffset(n, m, x, y int) {

	iterator := x*n + y
	for i := x * n; i < x*n+n; i++ {
		fmt.Print("{")
		for j := y; j < m+y; j++ {
			if j == y {
				fmt.Print(iterator)
			} else {
				fmt.Print(",", iterator)

			}
			iterator++
		}
		fmt.Println("},")
	}
	fmt.Println()
}
