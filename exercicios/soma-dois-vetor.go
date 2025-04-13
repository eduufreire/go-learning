package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3,3}, 6))
}


// essa solucao retornou com 22ms
func twoSum(nums []int, target int) []int {
	indices := [2]int{0, 0}
	for i := range len(nums){
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				indices[0] = i
				indices[1] = j
				break
			}
		}
	}
	return indices[:]
}

// essa retorna com 0ms
func twoSum2(nums []int, target int) []int {
    m := make(map[int]int, 0)
    for indx1, val := range nums {
        if indx2, exist := m[target - val]; exist {
            return []int{indx1, indx2}
        }
        m[val] = indx1
    }
    return []int{}
}