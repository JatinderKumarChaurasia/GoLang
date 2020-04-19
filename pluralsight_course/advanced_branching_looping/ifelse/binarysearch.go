package main

import "fmt"

type intSortedSlice []int

func main() {
	intList := intSortedSlice{2, 5, 7, 9, 11, 13, 17, 19, 23, 29, 39, 41, 91, 97, 2411, 2609, 4511, 5679, 9199}
	numberToSearch := 5679
	result := -1
	if numberToSearch > intList[0] && numberToSearch <= intList[len(intList)-1] {
		result = binarySearch(intList, 0, len(intList)-1, numberToSearch)
	}
	if result == -1 {
		fmt.Println("Number Not Found")
	} else {
		fmt.Printf("Number #%v found at %v Index", numberToSearch, result+1)
	}
}

func binarySearch(list intSortedSlice, left int, right int, search int) int {
	if right >= left {
		mid := (right + left) / 2
		if list[mid] == search {
			return mid
		}
		if list[mid] > search {
			return binarySearch(list, 0, mid-1, search)
		} else {
			return binarySearch(list, mid+1, right, search)
		}
	}
	return -1
}
