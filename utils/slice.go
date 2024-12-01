package utils

func CountOccurrences(slice []uint32, element uint32) uint32 {
	var count uint32 = 0
	for _, value := range slice {
		if value == element {
			count++
		}
	}
	return count
}
