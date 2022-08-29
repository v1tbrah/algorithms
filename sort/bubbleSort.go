package sort

func bubbleSort(input []int) {
	lenInput := len(input)
	if lenInput < 2 {
		return
	}
	for i := 0; i < len(input)-1; i++ {
		var flag bool
		for j := 0; j < len(input)-1; j++ {
			if input[j+1] < input[j] {
				input[j+1], input[j] = input[j], input[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return
}
