package common

func Sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func Product(nums []int) int {
	result := 1
	for _, num := range nums {
		result = result * num
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

func IndicesOf[T comparable](list []T, element T) []int {
	indices := []int{}

	for i, item := range list {
		if item == element {
			indices = append(indices, i)
		}
	}
	return indices
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

func Reverse[T comparable](list []T) []T {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}
