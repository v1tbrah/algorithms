package sort

func quickSort(input []int) {
	lenInput := len(input)
	if lenInput < 2 {
		return
	}
	if lenInput == 2 {
		if input[0] > input[1] {
			input[0], input[1] = input[1], input[0]
		}
		return
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
	quickSort(small)
	quickSort(big)
	for i := 0; i < len(small); i++ {
		input[i] = small[i]
	}
	input[len(small)] = supEl
	for i := 0; i < len(big); i++ {
		input[i+len(small)+1] = big[i]
	}
	return
}
