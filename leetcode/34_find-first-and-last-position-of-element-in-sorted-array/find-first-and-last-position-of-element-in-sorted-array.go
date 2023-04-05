package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	return []int{left(nums, target), right(nums, target)}
}

func left(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}

func right(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid
		}
	}

	if nums[r] == target {
		return r
	}
	if nums[l] == target {
		return l
	}
	return -1
}
