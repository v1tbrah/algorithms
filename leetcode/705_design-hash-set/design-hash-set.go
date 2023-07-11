package design_hash_set

// Task: https://leetcode.com/problems/design-hashset/

// Solution: https://leetcode.com/problems/design-hashset/submissions/992093781/

// tags:
// #Array
// #Hash Table
// #Linked List
// #Design
// #Hash Function

type MyHashSet struct {
	buckets [1000][]int
}

func Constructor() MyHashSet {
	return MyHashSet{
		buckets: [1000][]int{},
	}
}

func (this *MyHashSet) Add(key int) {
	if this == nil {
		return
	}

	if this.Contains(key) {
		return
	}

	bucketIdx := key % 1000

	this.buckets[bucketIdx] = append(this.buckets[bucketIdx], key)
}

func (this *MyHashSet) Remove(key int) {
	if this == nil {
		return
	}

	bucketIdx := key % 1000

	founded := this.buckets[bucketIdx]
	for i := 0; i < len(founded); i++ {
		if founded[i] == key {
			founded = append(founded[:i], founded[i+1:]...) // delete by idx from slice
			this.buckets[bucketIdx] = founded
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	if this == nil {
		return false
	}

	bucketIdx := key % 1000

	founded := this.buckets[bucketIdx]

	for i := 0; i < len(founded); i++ {
		if founded[i] == key {
			return true
		}
	}

	return false
}
