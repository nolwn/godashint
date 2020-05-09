package godashint

func findIndex(arr []int, predicate func(value int) bool, from int) int {
	for i := from; i < len(arr); i++ {
		if predicate(arr[i]) {
			return i
		}
	}

	return -1
}
