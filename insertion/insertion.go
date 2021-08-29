package insertion

func InsertionSort(A *[]int) *[]int {
	targetA := *A
	lenA := len(targetA)

	for j := 1; j < lenA; j++ {
		i := j - 1
		key := targetA[j]
		for i >= 0 && targetA[i] > key {
			targetA[i+1] = targetA[i]
			i = i - 1
		}
		targetA[i+1] = key
	}
	return &targetA
}
