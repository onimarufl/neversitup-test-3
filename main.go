package main

import "fmt"

func main() {
	slice_of_slices := make([][]int, 5)
	slice_of_slices[0] = []int{7}
	slice_of_slices[1] = []int{0}
	slice_of_slices[2] = []int{1, 1, 2}
	slice_of_slices[3] = []int{0, 1, 0, 1, 0}
	slice_of_slices[4] = []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}

	for _, value := range slice_of_slices {
		if len(value) == 1 {
			fmt.Println(fmt.Sprintf("should return %d, because it occurs %d time (which is odd).", value[0], len(value)))
			continue
		}

		duplicate := make(map[int]int)
		for _, v := range value {
			_, exist := duplicate[v]
			if exist {
				duplicate[v] += 1
			} else {
				duplicate[v] = 1
			}
		}

		for k, v := range duplicate {
			if v%2 != 0 {
				fmt.Println(fmt.Sprintf("should return %d, because it occurs %d time (which is odd).", k, v))
			}
		}
	}

}
