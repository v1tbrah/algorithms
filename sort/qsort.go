package sort

func Qsort(input []int) []int {
	lenInput := len(input)
	if lenInput < 2 {
		return input
	}
	if lenInput == 2 {
		if input[0] > input[1] {
			input[0], input[1] = input[1], input[0]
		}
		return input
	}
	idxSupEl := lenInput / 2
	supEl := input[idxSupEl]
	small := []int{}
	big := []int{}
	for i, v := range input {
		if i == idxSupEl {
			continue
		}
		if v < supEl {
			small = append(small, v)
		} else {
			big = append(big, v)
		}
	}
	sortedSmall := Qsort(small)
	sortedBig := Qsort(big)
	for i := 0; i < len(sortedSmall); i++ {
		input[i] = sortedSmall[i]
	}
	input[len(sortedSmall)] = supEl
	for i := 0; i < len(sortedBig); i++ {
		input[i+len(sortedSmall)+1] = sortedBig[i]
	}
	return input
}
