package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, el := range nums {
		for j, elm := range nums {
			if el+elm == target && j > i {
				return []int{i , j}
			}
		}
	}
	// now it's OUTSIDE the loop
	return []int{0}
}


func main()  {
	nums := []int{2,7,11,15} 
	target := 9
	fmt.Println(twoSum(nums, target))
}