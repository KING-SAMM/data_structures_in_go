/*-----------------------------------------------------
| ARRAY IMPLEMENTATION OF THE MAX-HEAP DATA STRUCTURE
------------------------------------------------------*/
package main

import "fmt"

// Define a struct for the minHeap
type maxHeap struct {
	storage []int
	size int
}

// Create a new max heap
func createNewHeap() *maxHeap {
	return &maxHeap{
		storage: make([]int, 0),
		size: 0,
	}
}

// Define functions to get parent, leftChild and rightChild indices
func (mxh *maxHeap) GetParentIndex(currentIndex int) int {
	return (currentIndex - 1)/ 2
}

func (mxh *maxHeap) GetLeftChildIndex(currentIndex int) int {
	return 2 * currentIndex + 1
}

func (mxh *maxHeap) GetRightChildIndex(currentIndex int) int {
	return 2 * currentIndex + 2
}

// Check if parent, leftChild or rightChild exist
func (mxh *maxHeap) HasParent(currentIndex int) bool {
	return mxh.GetParentIndex(currentIndex) >= 0 
}

func (mxh *maxHeap) HasLeftChild(currentIndex int) bool {
	return mxh.GetLeftChildIndex(currentIndex) < mxh.size
}

func (mxh *maxHeap) HasRightChild(currentIndex int) bool {
	return mxh.GetRightChildIndex(currentIndex) < mxh.size
}

// Get actual parent, leftChild or rightChild node data
func (mxh *maxHeap) Parent(currentIndex int) int {
	return mxh.storage[mxh.GetParentIndex(currentIndex)]
}

func (mxh *maxHeap) LeftChild(currentIndex int) int {
	return mxh.storage[mxh.GetLeftChildIndex(currentIndex)]
}

func (mxh *maxHeap) RightChild(currentIndex int) int {
	return mxh.storage[mxh.GetRightChildIndex(currentIndex)]
}

func (mxh *maxHeap) Swap(index1 int, index2 int) {
	temp := mxh.storage[index1]
	mxh.storage[index1] = mxh.storage[index2]
	mxh.storage[index2] = temp
}

// Insert data into the heap
func (mxh *maxHeap) Insert(data int) {
	mxh.storage = append(mxh.storage, data)
	mxh.size += 1
	mxh.HeapifyUp()
}

// We advance up the tree by perform heapify-up operation
func (mxh *maxHeap) HeapifyUp() {
	// Starting with the index of the inserted node(data)
	currentIndex := mxh.size - 1
	// So long as a parent node exists, and its value is lLESS than that
	// of the current node, we swap parent node with current node
	for mxh.HasParent(currentIndex) && mxh.Parent(currentIndex) < mxh.storage[currentIndex] {
		mxh.Swap(mxh.GetParentIndex(currentIndex), currentIndex)
		// Reassign current index to reflect swap
		currentIndex = mxh.GetParentIndex(currentIndex)
	}
}

// Print the heap
func (mxh *maxHeap) Print() {
	fmt.Println(mxh.storage)
}

func main() {
	max_heap := createNewHeap()

	max_heap.Insert(10)
	max_heap.Insert(20)
	max_heap.Insert(5)
	max_heap.Insert(30)
	max_heap.Insert(15)
	max_heap.Insert(3)
	max_heap.Insert(0)
	// fmt.Println(max_heap)
	max_heap.Print()
	fmt.Println(max_heap.size)
}

