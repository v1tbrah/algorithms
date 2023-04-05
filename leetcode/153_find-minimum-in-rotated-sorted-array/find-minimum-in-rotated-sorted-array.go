package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	l, r, pivot := 0, len(nums)-1, 0
	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] > nums[mid+1] {
			pivot = mid + 1
			break
		}

		if nums[mid] < nums[mid-1] {
			pivot = mid
			break
		}

		if nums[0] > nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return nums[pivot]
}
