/*-----------------------------------------------------
| ARRAY IMPLEMENTATION OF THE MAX-HEAP DATA STRUCTURE
------------------------------------------------------*/
package main

import (
	"fmt"
	"errors"
)

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

// Remove the maximum data from the heap
func (mxh *maxHeap) RemoveMax() int {
	// Display an error if heap is empty 
	if mxh.size == 0 {
		fmt.Println(errors.New("Empty heap!"))
	}
	// Get the maximum element/value which should be at the root
	data := mxh.storage[0]
	// Replace it with last node/element in the heap
	mxh.storage[0] = mxh.storage[mxh.size - 1]
	// Update heap (slice)
	mxh.storage = mxh.storage[:mxh.size - 1]
	// Update (reduce) the heap size
	mxh.size -= 1
	// place new node at root in right position by heapifying down
	mxh.HeapifyDown()
	// Return the removed largest node
	return data
}

// Heapify down
func (mxh *maxHeap) HeapifyDown() {
	// Starting at the root position
	currentIndex := 0
	// so long as left child exists
	for mxh.HasLeftChild(currentIndex) {
		// grab its index, set it to largerChildIndex
		largerChildIndex := mxh.GetLeftChildIndex(currentIndex)
		// Check if right child exists and if it is greater than left child
		if mxh.HasRightChild(currentIndex) && mxh.RightChild(currentIndex) > mxh.LeftChild(currentIndex) {
			// If true reassign largerChildIndex to right child's index
			largerChildIndex = mxh.GetRightChildIndex(currentIndex)
		}
		// Check if current node is greater than selected (larger )child node 
		if mxh.storage[currentIndex] > mxh.storage[largerChildIndex] {
			// if true break out of loop
			break
		} else {
			// else swap their positions
			mxh.Swap(currentIndex, largerChildIndex)
			// Move new position to node at largerChildIndex
			currentIndex = largerChildIndex 
		}
	}

}


// Print the heap
func (mxh *maxHeap) Print() {
	fmt.Println(mxh.storage)
}

func main() {
	max_heap := createNewHeap()

	// Insert a number of nodes
	max_heap.Insert(10)
	max_heap.Insert(20)
	max_heap.Insert(5)
	max_heap.Insert(8)
	max_heap.Insert(0)
	max_heap.Insert(15)
	max_heap.Insert(30)

	// Remove max and print in succession
	max_heap.Print()
	fmt.Println("Size:", max_heap.size)
	fmt.Println("Max:", max_heap.RemoveMax())
	max_heap.Print()
	fmt.Println("Size:", max_heap.size)
	fmt.Println("Max:", max_heap.RemoveMax())
	max_heap.Print()
	fmt.Println("Size:", max_heap.size)

	// Output:
	// [30 10 20 8 0 5 15]
	// Size: 7
	// Max: 30
	// [20 10 15 8 0 5]
	// Size: 6
	// Max: 20
	// [15 10 5 8 0]
	// Size: 5
}

