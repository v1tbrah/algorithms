package insert_delete_getrandom_o1

import "math/rand"

// Task: https://leetcode.com/problems/insert-delete-getrandom-o1/

// Solution: https://leetcode.com/problems/insert-delete-getrandom-o1/submissions/1067096020/

// tags:
// #Array
// #Hash Table
// #Math
// #Design
// #Randomized

type RandomizedSet struct {
	arr      []int
	valToIdx map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr:      make([]int, 0),
		valToIdx: make(map[int]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToIdx[val]; ok {
		return false
	}

	this.arr = append(this.arr, val)
	this.valToIdx[val] = len(this.arr) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.valToIdx[val]
	if !ok {
		return false
	}

	// swap with last
	this.arr[idx] = this.arr[len(this.arr)-1]
	// save new idx for swapped element
	this.valToIdx[this.arr[idx]] = idx
	// cut last
	this.arr = this.arr[:len(this.arr)-1]
	// delete last val:idx from map
	delete(this.valToIdx, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
