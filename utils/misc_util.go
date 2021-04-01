package utils

// Contains return bool
// contains checks if a string is present in a slice
//ref: https://play.golang.org/p/Qg_uv_inCek
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
