package utils

func CheckStringInArray(str string, arr []string) bool {
	for _, value := range arr {
		if value == str {
			return true
		}
	}
	return false
}
