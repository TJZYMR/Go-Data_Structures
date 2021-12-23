package main

import "fmt"

// Enqueue: Adds an item to the queue. If the queue is full, then it is said to be an Overflow condition – Time Complexity : O(1)
// Dequeue: Removes an item from the queue. The items are popped in the same order in which they are pushed. If the queue is empty, then it is said to be an Underflow condition – Time Complexity : O(1)
// Front: Get the front item from queue – Time Complexity : O(1)
// Rear: Get the last item from queue – Time Complexity : O(1)
type Queue struct {
	Items []int
	Rear  int //holds the index of the rear element
	Front int //holds the index of the front element
}

func (Q *Queue) Enqueue(Item int) {
	if len(Q.Items) == 0 {
		Q.Items = append(Q.Items, Item)
		Q.Rear = 0
		Q.Front = 0
	} else {
		Q.Items = append(Q.Items, Item)
		Q.Front = len(Q.Items) - 1
	}
}
func (Q *Queue) Dequeue() {
	if len(Q.Items) == 0 {
		fmt.Println("Queue is empty")
		Q.Rear = -1

	} else if len(Q.Items) == 1 {
		Q.Items = Q.Items[:len(Q.Items)-1]
		Q.Rear = -1
		Q.Front = len(Q.Items) - 1
	} else {
		Q.Items = Q.Items[:len(Q.Items)-1]
		Q.Rear = 0
		Q.Front = len(Q.Items) - 1
	}
}
func main() {
	q := []int{} //initially empty stack
	queue := &Queue{Items: q, Rear: 0, Front: 0}
	queue.Enqueue(11)
	queue.Enqueue(12)
	queue.Enqueue(13)
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Enqueue(11)
	queue.Enqueue(12)
	queue.Enqueue(13)
	queue.Dequeue()

	queue.Enqueue(13)
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	fmt.Println("Queue Status:==>", queue)
}
