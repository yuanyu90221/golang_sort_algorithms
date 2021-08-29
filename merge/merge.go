package merge

func MergeSort(A *[]int, p int, r int) *[]int {
	targetA := (*A)
	if p < r {
		q := (p + r) / 2
		MergeSort(A, p, q)
		MergeSort(A, q+1, r)
		Merge(A, p, q, r)
	}
	return &targetA
}
func Merge(A *[]int, p int, q int, r int) *[]int {
	targetA := (*A)
	L := make([]int, q-p+1)
	R := make([]int, r-q)
	copy(L, targetA[p:q+1])
	copy(R, targetA[q+1:r+1])
	i := 0
	j := 0
	for k := p; k <= r; k++ {
		targetA[k], i, j = FindSmallElement(&L, &R, i, j)
	}
	return &targetA
}
func FindSmallElement(L *[]int, R *[]int, i int, j int) (int, int, int) {
	if i >= len(*L) {
		return (*R)[j], i, j + 1
	}
	if j >= len(*R) {
		return (*L)[i], i + 1, j
	}
	if (*L)[i] <= (*R)[j] {
		return (*L)[i], i + 1, j
	} else {
		return (*R)[j], i, j + 1
	}
}
