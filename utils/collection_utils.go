package utils

func Contains(intSlice []int, searchInt int) bool {
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}

func Remove(intSlice []int, removeInt int) []int{
	intSlice = append(intSlice[:removeInt], intSlice[removeInt+1:]...)
	return intSlice
}


