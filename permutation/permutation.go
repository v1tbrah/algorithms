package permutation

func fillPermutations(s []byte, i, n int, permutations *[][]byte) {

	if i == n-1 {
		permutation := make([]byte, len(s))
		copy(permutation, s)
		*permutations = append(*permutations, permutation)
		return
	}

	for j := i; j < n; j++ {
		s[i], s[j] = s[j], s[i]
		fillPermutations(s, i+1, n, permutations)
		s[i], s[j] = s[j], s[i]
	}
}
