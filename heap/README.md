# MaxHeap/MinHeap
Both Heaps are concurrent safety

## Methods
* Push
* Pop (delete top)
* GetSize
* GetPeek (get top)
* String (only for debug)

## MinHeap
* Push O(logN)
* Pop (delete top) O(logN)
* GetPeek O(1)

Example
```
   1
2     3
```

## MaxHeap
* Push O(logN)
* Pop (delete top) O(logN)
* GetPeek O(1)

```
   3
1     2
```

## Implementation
```
vertex              ║ index                  ║
╠═══════════════════╬════════════════════════╣
║ root              ║ 0                      ║
║ current           ║ i                      ║
║ parent            ║ (i - 1) / 2            ║
║ left child        ║ 2*i + 1                ║
║ right child       ║ 2*i + 2                ║
║ the last non-leaf ║ (array length / 2) - 1 ║
```