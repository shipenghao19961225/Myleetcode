package leetcode

func toLowerCase(str string) string {
	byteArray := []byte(str)
	for i := 0; i < len(byteArray); i++ {
		if byteArray[i] >= 'A' && byteArray[i] <= 'Z' {
			byteArray[i] = 'a' + byteArray[i] - 'A'
		}
	}
	return string(byteArray)
}
