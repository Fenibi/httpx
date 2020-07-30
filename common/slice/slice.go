package slice

// IntSliceContains check if a slice contains the specified int value
func IntSliceContains(sl []int, v int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}
