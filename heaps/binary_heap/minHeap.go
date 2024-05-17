/*-----------------------------------------------------
| ARRAY IMPLEMENTATION OF THE MIN-HEAP DATA STRUCTURE
------------------------------------------------------*/
package main

import (
	"fmt"
	"errors"
)

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
	mh.HeapifyUp()
}

// Perform heapify-up operation ()
func (mh *minHeap) HeapifyUp() {
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

// Remove the minimum data from the heap
func (mh *minHeap) RemoveMin() int{
	// Return an error if heap is empty
	if mh.size == 0 {
        fmt.Println(errors.New("Empty heap!"))
    }
	// Get the minimum element/value which should be at the root
	data := mh.storage[0]
	// Replace the element at the root with the last element in the heap
	mh.storage[0] = mh.storage[mh.size - 1]
	// Update the slice
	mh.storage = mh.storage[: mh.size - 1]
	// Reduce the heap size
	mh.size -= 1
	// Heapify down
	mh.HeapifyDown()
	// Return the removed data
    return data
}

// Heapify down
func (mh *minHeap) HeapifyDown() {
	// Starting from the root node position
	currentIndex := 0
	// So long as left child exists
	for mh.HasLeftChild(currentIndex) {
		// grab its index
		smallerChildIndex := mh.GetLeftChildIndex(currentIndex)
		// Check if Right Child exists and if it's smaller than left child 
		if mh.HasRightChild(currentIndex) && mh.RightChild(currentIndex) < mh.LeftChild(currentIndex) {
			// If true reassign smallerChildIndex to RightChild's index
			smallerChildIndex = mh.GetRightChildIndex(currentIndex)
		}
		// Check if current node is less than selected (smaller) child node
		if mh.storage[currentIndex] < mh.storage[smallerChildIndex] {
			// Break out of loop if true
			break
		} else {
			// Else swap their positions
			mh.Swap(currentIndex, smallerChildIndex)
			// Move new position to node at smallerChildIndex
			currentIndex = smallerChildIndex
		}
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
	min_heap.Insert(8)
	min_heap.Insert(0)
	min_heap.Insert(15)
	min_heap.Insert(30)

	// fmt.Println(min_heap)
	min_heap.Print()
	fmt.Println(min_heap.size)
	fmt.Println("Min:", min_heap.RemoveMin())
	min_heap.Print()
	fmt.Println(min_heap.size)
	fmt.Println("Min:", min_heap.RemoveMin())
	min_heap.Print()
	fmt.Println(min_heap.size)
 }