# BINARY HEAPS

## Types

1. Min Heap
2. Max Heap

## Properties

1. A parent node must have at most two children
2. It must be a complete tree. It must be filled from left to right, and every level must be full, with the exception of the last level which need NOT be full.
                   
### A Complete Tree 

All levels are full

<img src="../assets/complete_tree1.png" alt="Complete tree 1" width="400">
  
  
### A Complete Tree 

First two levels are full, last level not full but filled from left to right

<img src="../assets/complete_tree2.png" alt="Complete tree 2" width="400">           
 

### An Incomplete Tree 

Level 1 is not full before moving to level 2

<img src="../assets/incomplete_tree1.png" alt="Incomplete tree 1" width="400">   
  
  
### An incomplete Tree 

Level 2 is filled from the right instead of left

<img src="../assets/incomplete_tree2.png" alt="Incomplete tree 2" width="400">   
   
3. Min-heap: Every parent's key (node value) must be smaller than its children nodes'. This ensures the parent key is the smallest within the heap. 
Max-heap: Every parent's key (node value) must be larger than its children nodes'. This ensures the parent key is the largest within the heap.


## Representing A Heap in an Array

To do this we use the following formulae:
  
- Calculate parent Index: Floor of (current index - 1)/ 2
- Calculate left-child Index: 2 * parentIndex + 1
- Calculate Right-child Index: 2 * parentIndex + 2

See the Min-Heap [code here](./binary_heap/minHeap.go), and Max-Heap [here](./binary_heap/maxHeap.go)

## Steps
1. Define a struct for the minHeap
```go
type minHeap struct{
	storage []int
	size int
}
```
2. Create a new minHeap
```go
func newHeap() *minHeap {
	return &minHeap{
		storage: make([]int, 0),
		size: 0,
	} 
}
```
We should be able to:
3. Get parent, left-child and right-child positions 
4. Determine if the current node has parent, left-child or right-child, and return a boolean
5. Get actual data from parent, left-child and right-child 
6. Swap elements at given indices
7. Insert data into the heap
8. Perform heapify-up operation
9. Remove the minimum data from the heap
10. Heapify down
11. Print the heap

See the Min-Heap [code here](./binary_heap/minHeap.go), and Max-Heap [here](./binary_heap/maxHeap.go)