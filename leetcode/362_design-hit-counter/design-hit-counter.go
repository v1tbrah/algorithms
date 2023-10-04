package design_hit_counter

// Task: https://leetcode.com/problems/design-hit-counter/

// Solution: https://leetcode.com/problems/design-hit-counter/submissions/1066966831/

// tags:
// #Array
// #Hash Table
// #Design

type HitCounter struct {
	heats map[int]int
}

func Constructor() HitCounter {
	return HitCounter{
		heats: make(map[int]int),
	}
}

// Time: O(1)
func (this *HitCounter) Hit(timestamp int) {
	this.heats[timestamp]++
}

// Time: O(n)
func (this *HitCounter) GetHits(timestamp int) int {
	var result int
	for i, v := range this.heats {
		if timestamp-300 < i {
			result += v
		} else {
			delete(this.heats, i)
		}
	}

	return result
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
