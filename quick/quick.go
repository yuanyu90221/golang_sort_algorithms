package quick

func QuickSort(A *[]int, p int, r int) *[]int {
	targetA := (*A)
	if p < r {
		q := Partition(A, p, r)
		QuickSort(A, p, q-1)
		QuickSort(A, q+1, r)
	}
	return &targetA
}
func Partition(A *[]int, p int, r int) int {
	targetA := (*A)
	x := targetA[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if targetA[j] <= x {
			i = i + 1
			targetA[i], targetA[j] = targetA[j], targetA[i]
		}
	}
	targetA[i+1], targetA[r] = targetA[r], targetA[i+1]
	return i + 1
}
