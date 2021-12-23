package main

// Enqueue: Adds an item to the queue. If the queue is full, then it is said to be an Overflow condition – Time Complexity : O(1)
// Dequeue: Removes an item from the queue. The items are popped in the same order in which they are pushed. If the queue is empty, then it is said to be an Underflow condition – Time Complexity : O(1)
// Front: Get the front item from queue – Time Complexity : O(1)
// Rear: Get the last item from queue – Time Complexity : O(1)
type Queue struct {
	Items []int
	Rear  *int //holds the index of the rear element
	Front *int //holds the index of the front element
}

func (Q *Queue) Enqueue(Item int) {

}
func main() {

}
