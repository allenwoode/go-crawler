package algo

var lastOccurred = make([]int, 0xffff)

func LenOfNoRepeatingString(str string) int {
	//lastOccurred := make(map[rune]int)
	//lastOccurred := make([]int, 0xffff)
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	start := 0
	maxLength := 0

	for i, ch := range []rune(str) {
		if last := lastOccurred[ch]; last != -1 && last >= start {
			start = last + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
