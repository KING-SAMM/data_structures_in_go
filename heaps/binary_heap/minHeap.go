/*-----------------------------------------------------
| ARRAY IMPLEMENTATION OF THE MIN-HEAP DATA STRUCTURE
------------------------------------------------------*/
package main

import ("fmt")

// Define a struct for the minHeap
type minHeap struct{
	storage []int
	size int
}

// Create a new minHeap
func newHeap() *minHeap {
	return &minHeap{
		storage: make([]int, 0),
		size: 0,
	} 
}

// Get parent, left-child and right-child positions 
func (mh *minHeap) GetParentIndex(currentIndex int) int {
	// Integer division discards fractional part
	return (currentIndex - 1)/2

}

func (mh *minHeap) GetLeftChildIndex(curerntIndex int) int {
	return 2 * curerntIndex + 1
}

func (mh *minHeap) GetRightChildIndex(curerntIndex int) int {
	return 2 * curerntIndex + 2
}

// Determine if node has parent, left-child or right-child. Returns a boolean
func (mh *minHeap) HasParent(currentIndex int) bool {
	return mh.GetParentIndex(currentIndex) >= 0
}

func (mh *minHeap) HasLeftChild(currentIndex int) bool {
	return mh.GetLeftChildIndex(currentIndex) < mh.size
}

func (mh *minHeap) HasRightChild(currentIndex int) bool {
	return mh.GetRightChildIndex(currentIndex) < mh.size
}

// Get actual data from parent, left-child and right-child 
func (mh *minHeap) Parent(currentIndex int) int {
	return mh.storage[mh.GetParentIndex(currentIndex)]
}

func (mh *minHeap) LeftChild(currentIndex int) int {
	return mh.storage[mh.GetLeftChildIndex(currentIndex)]
}

func (mh *minHeap) RightChild(currentIndex int) int {
	return mh.storage[mh.GetRightChildIndex(currentIndex)]
}

// Swap elements at given indices
func (mh *minHeap) Swap(index1, index2 int) {
	temp := mh.storage[index1]
	mh.storage[index1] = mh.storage[index2]
	mh.storage[index2] = temp
}

// Insert data into the heap
func (mh *minHeap) Insert(data int) {
	// First insert data in the last available position
	mh.storage = append(mh.storage, data)
	mh.size += 1
	// Make sure the data ends up in the correct position
	mh.heapifyUp()
}

// Perform heapify-up operation
func (mh *minHeap) heapifyUp() {
	// starting from index of our newly inserted node
	currentIndex := mh.size - 1
	// keep walking up the Tree, heapifying as we go
	// so long as we have a parent node, and the value of the parent node
	// is GREATER than the current node
	for mh.HasParent(currentIndex) && mh.Parent(currentIndex) > mh.storage[currentIndex] {
		// As long as the above is true we swap to maintain the minHeap
		// property that the parent node must be less than the child
		mh.Swap(mh.GetParentIndex(currentIndex), currentIndex)
		// After swapping we continue advancing up the tree
		currentIndex = mh.GetParentIndex(currentIndex)
	}
}

// Print the heap
func (mh *minHeap) Print() {
	fmt.Println(mh.storage)
	// return mh.storage 
}

func main() {
	min_heap := newHeap()

	min_heap.Insert(10)
	min_heap.Insert(20)
	min_heap.Insert(5)
	min_heap.Insert(30)
	min_heap.Insert(15)
	min_heap.Insert(3)
	min_heap.Insert(0)
	// fmt.Println(min_heap)
	min_heap.Print()
	fmt.Println(min_heap.size)
 }