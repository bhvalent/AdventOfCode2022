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

func Distinct[T comparable](items []T) []T {
	keys := make(map[T]bool)
	list := []T{}

	for _, item := range items {
		if value := keys[item]; !value {
			keys[item] = true
			list = append(list, item)
		}
	}

	return list
}
