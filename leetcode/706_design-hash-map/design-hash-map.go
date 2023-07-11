package design_hash_map

// Task: https://leetcode.com/problems/design-hashmap/

// Solution: https://leetcode.com/problems/design-hashmap/submissions/992102963/

// tags:
// #Array
// #Hash Table
// #Linked List
// #Design
// #Hash Function

type MyHashMap struct {
	buckets [1000][]keyValue
}

type keyValue struct {
	key   *int
	value int
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets: [1000][]keyValue{},
	}
}

func (this *MyHashMap) Put(key int, value int) {
	if this == nil {
		return
	}

	bucketIdx := key % 1000

	founded := this.buckets[bucketIdx]

	for i := 0; i < len(founded); i++ {
		if founded[i].key != nil && *founded[i].key == key {
			founded[i].value = value
			return
		}
	}

	this.buckets[bucketIdx] = append(this.buckets[bucketIdx], keyValue{&key, value})
}

func (this *MyHashMap) Get(key int) int {
	if this == nil {
		return 0
	}

	bucketIdx := key % 1000

	founded := this.buckets[bucketIdx]

	for i := 0; i < len(founded); i++ {
		if founded[i].key != nil && *founded[i].key == key {
			return founded[i].value
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	if this == nil {
		return
	}

	bucketIdx := key % 1000

	founded := this.buckets[bucketIdx]

	for i := 0; i < len(founded); i++ {
		if founded[i].key != nil && *founded[i].key == key {
			founded = append(founded[:i], founded[i+1:]...) // delete by idx from slice
			this.buckets[bucketIdx] = founded
			return
		}
	}
}
