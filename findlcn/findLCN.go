package findlcn

func LongestConsecutive(nums []int) int {
	var longesetAccum int = 0
	valueSet := NewSet(nums...)
	var values []int = valueSet.Values()
	for _, num := range values {
		if !valueSet.Contains(num - 1) {
			currentNum := num
			currentAccum := 1
			for valueSet.Contains(currentNum + 1) {
				currentNum += 1
				currentAccum += 1
			}
			if currentAccum > longesetAccum {
				longesetAccum = currentAccum
			}
		}
	}
	return longesetAccum
}
