package consistentHashing

type uints []uint32

// return the len of slice
func (x uints) Len() int {
	return len(x)
}

// compare element of slice
func (x uints) Less(i, j int) bool {
	return x[i] < x[j]
}

// swap element of slice
func (x uints) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
