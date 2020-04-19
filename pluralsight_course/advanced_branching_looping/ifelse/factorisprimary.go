package main

import "fmt"

type intSlice []int

func main() {
	intList := intSlice{24, 27, 224, 123, 484, 364, 999, 789, 764, 124, 500, 2345, 2047083082, 239, 111}
	factors := make(map[int]intSlice)

	for _, i := range intList {
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				factors[i] = append(factors[i], j)
			}
		}
	}

	for key, val := range factors {
		fmt.Printf("The Factors of %d is %v\n", key, val)
	}

	fmt.Print("\n Primary Numbers are : \t")
	for key, val := range factors {
		if len(val) != 2 {
			continue
		}
		fmt.Printf("%d\t", key)
	}
}
