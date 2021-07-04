package sort

//BubbleSort

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

//quickSort

func partition(arr []int, low, high int) int {

	p := arr[high]

	for j := low; j < high; j++ {
		if arr[j] < p {
			arr[j], arr[low] = arr[low], arr[j]
			low++
		}
	}

	arr[low], arr[high] = arr[high], arr[low]
	return low
}

func quickSort(arr []int, low, high int) []int {

	if low > high {
		return arr
	}

	p := partition(arr, low, high)
	quickSort(arr, low, p-1)
	quickSort(arr, p+1, high)
	return arr
}

//MergeSort

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	first := MergeSort(array[:len(array)/2])
	second := MergeSort(array[len(array)/2:])
	return Merge(first, second)
}

func Merge(a []int, b []int) []int {
	var final []int
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			final = append(final, b[j])
			j++
		} else {
			final = append(final, a[i])
			i++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}
