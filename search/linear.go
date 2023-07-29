package search

// Linear search algorithm.
func Linear(arr []int, target int) (int, error) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i, nil
		}
	}
	return -1, ErrNotFound
}
