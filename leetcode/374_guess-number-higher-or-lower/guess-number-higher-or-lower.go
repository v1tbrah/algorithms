package guess_number_higher_or_lower

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		mid := r - (r-l)/2
		switch {
		case guess(mid) == -1: // higher picked
			r = mid - 1
		case guess(mid) == 1: // lower picked
			l = mid + 1
		case guess(mid) == 0:
			return mid
		default:
			panic("unexpected 'guess' func behaviour")
		}

	}
	return -1
}

// @return
//
//	  -1 if num is higher than the picked number
//		  1 if num is lower than the picked number
//		  otherwise return 0
func guess(val int) int {
	return -1 // dummy
}
