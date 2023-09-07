package insert_delete_getrandom_o1

import "math/rand"

// Task: https://leetcode.com/problems/insert-delete-getrandom-o1/

// Solution: https://leetcode.com/problems/insert-delete-getrandom-o1/submissions/1043246844/

// tags:
// #Array
// #Hash Table
// #Math
// #Design
// #Randomized

type RandomizedSet struct {
	data     []int
	valToIdx map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data:     make([]int, 0),
		valToIdx: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToIdx[val]; ok {
		return false
	}

	this.data = append(this.data, val)
	this.valToIdx[val] = len(this.data) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.valToIdx[val]
	if !ok {
		return false
	}

	// swap last and this
	this.data[idx], this.data[len(this.data)-1] = this.data[len(this.data)-1], this.data[idx]
	// update idx for last
	this.valToIdx[this.data[idx]] = idx
	// cut last element
	this.data = this.data[:len(this.data)-1]

	delete(this.valToIdx, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
