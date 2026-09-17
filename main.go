package main

import "fmt"

// Two Sum homework

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		need := target - num

		if j, ok := seen[need]; ok {
			return []int{j, i}
		}

		seen[num] = i
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	fmt.Println(result)
}
