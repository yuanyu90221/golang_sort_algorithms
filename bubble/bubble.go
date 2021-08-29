package bubble

func BubbleSort(A *[]int) *[]int {
	targetA := (*A)
	lenA := len(targetA)
	rangeA := lenA - 1
	for i := 0; i < lenA; i++ {
		for j := rangeA; j >= i+1; j-- {
			if targetA[j] < targetA[j-1] {
				targetA[j], targetA[j-1] = targetA[j-1], targetA[j]
			}
		}
	}
	return &targetA
}
