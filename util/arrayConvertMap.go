package util

// ArrayIntConvertMap 将数据转换为map
// 注意：此方法只能转换int类型数组
func ArrayIntConvertMap(array []int64) map[int64]int64 {
	data := make(map[int64]int64, len(array))
	for _, v := range array {
		data[v] = 1
	}
	return data
}
