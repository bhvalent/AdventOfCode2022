package common

func Sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func IndexOf[T comparable](list []T, element T) int {
	for i, item := range list {
		if item == element {
			return i
		}
	}
	return -1
}
