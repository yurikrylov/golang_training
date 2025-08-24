package utils

func Find(arr []int, k int) []int {
	m := make(map[int]int)
	for i, a := range arr {
		if j, ok := m[k-a]; ok {
			return []int{i, j}
		}
		m[a] = i
	}
	return nil
}
