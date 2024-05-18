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
  
- Calculate parent Index: Floor of `(currentIndex - 1)/ 2`
- Calculate left-child Index: `2 * currentIndex + 1`
- Calculate Right-child Index: `2 * currentIndex + 2`

See the Min-Heap [code here](./binary_heap/minHeap.go), and Max-Heap [here](./binary_heap/maxHeap.go)

## Steps

1. Define a struct for the minHeap
```go
type minHeap struct{
	storage []int
	size int
}
```
2. Create a new minHeap with initial values
```go
func newHeap() *minHeap {
	return &minHeap{
		storage: make([]int, 0),
		size: 0,
	} 
}
```
We should be able to:

3. Get parent, left-child and right-child positions (indices)
4. Determine if the current node has parent, left-child or right-child, and return a boolean
5. Get actual data from parent, left-child and right-child nodes
6. Swap elements/nodes at given indices
7. Insert data (node) into the heap
8. Perform heapify-up operation
9. Remove the minimum data from the heap
10. Heapify down
11. Print the heap

## Run the code

Instantiate the heap
```go
    min_heap := newHeap()
```

Insert some nodes
```go
    min_heap.Insert(10)
    min_heap.Insert(20)
    min_heap.Insert(5)
    min_heap.Insert(8)
    min_heap.Insert(0)
    min_heap.Insert(15)
    min_heap.Insert(30)
```

Execute the removeHeap method in succession and print out the heap and size as you go
```go
    min_heap.Print()
    fmt.Println("Size:", min_heap.size)
    fmt.Println("Min:", min_heap.RemoveMin())
    min_heap.Print()
    fmt.Println("Size:", min_heap.size)
    fmt.Println("Min:", min_heap.RemoveMin())
    min_heap.Print()
    fmt.Println("Size:", min_heap.size)
```

Output:
```bash
    [0 5 10 20 8 15 30]
    Size: 7
    Min: 0
    [5 8 10 20 30 15]
    Size: 6
    Min:  5
    [8 15 10 20 30]
    Size: 5
```

See the Min-Heap [code here](./binary_heap/minHeap.go), and Max-Heap [here](./binary_heap/maxHeap.go)