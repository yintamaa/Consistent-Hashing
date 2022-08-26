package consistentHashing

type units []uint32

//返回切片长度
func (x units) Len() int {
	return len(x)
}

//比较两个值的大小
func (x units) Less(i, j int) bool {
	return x[i] < x[j]
}

//切片中值交换
func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
