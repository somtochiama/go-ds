package sorting

import "fmt"

func MergeSort(arr []int) []int{
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr)/2
	fmt.Println(mid)
	firstHalf := MergeSort(arr[:mid])
	secondHalf := MergeSort(arr[mid:])

	return merge(firstHalf, secondHalf)
}

func merge(arr1 []int, arr2 []int) []int{
	if len(arr1) == 0 {
		return arr2
	}

	if len(arr2) == 0 {
		return arr1
	}

	i := 0
	j := 0
	var ret []int

	for i < len(arr1) || j < len(arr2) {
		if i == len(arr1) {
			return append(ret, arr2[j:]...)
		}

		if j == len(arr2) {
			return append(ret, arr1[i:]...)
		}

		if arr1[i] < arr2[j] {
			ret = append(ret, arr1[i])
			i++
		} else {
			ret = append(ret, arr2[j])
			j++
		}
	}

	return ret
}