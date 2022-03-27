package sort

func Merge(left, right []int) []int {

	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {

		if len(left) == 0 {
			return append(result, right...)
		}

		if len(right) == 0 {
			return append(result, left...)
		}

		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func MergeSort(arr []int) []int {

	// 값이 0 혹은 1이면 arr을 그대로 반환
	if len(arr) <= 1 {
		return arr
	}

	// 배열의 중간 찾기
	middle := len(arr) / 2

	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	return Merge(left, right)
}
