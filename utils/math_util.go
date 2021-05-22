package utils

//MaxInt64 return maximum value
func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

//MaxInt return maximum value
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//MinInt return minimum value
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
