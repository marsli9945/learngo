package main

func twoSum(nums []int, target int) []int {
	var result = []int{0, 0}
	for _, v := range nums {
		for _, vv := range nums {
			if v+vv == target {
				result[0] = v
				result[1] = vv
				return result
			}
		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	sum := twoSum(nums, 9)
	println(sum[0], sum[1])
}
