# MaxHeap/MinHeap
Both Heaps are concurrent safety

## Methods
* Add
* Pop (delete top)
* GetSize
* GetPeek (get top)
* String (only for debug)

## MinHeap
* Insert O(logN)
* Pop (delete top) O(logN)
* Get min O(1)

Example
```
   1
2     3
```

## MaxHeap
* Insert O(logN)
* Pop (delete top) O(logN)
* Get max O(1)

```
   3
1     2
```

## Implementation
The main idea:
1. index of the parent node of any node is `index of the node / 2`
2. index of the left child node is `index of the node * 2`
3. index of the right child node is `index of the node * 2 + 1`
4. node is a leaf node, when `idx > h.realSize/2`