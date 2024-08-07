package main

import "fmt"

// Queue represents a queue data structure
type Queue struct {
	result []int
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(element int) {
	q.result = append(q.result, element)
}

// Dequeue removes and returns the first element of the queue
func (q *Queue) Dequeue() (int, bool) {
	if len(q.result) == 0 {
		return 0, false
	}
	element := (q.result)[0]
	q.result = (q.result)[1:]
	return element, true
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.result) == 0
}

// This methode return true if the element in the Queue otherwise false
func (q *Queue) contains(isthere int) bool {
	for i := 0; i < len(q.result); i++ {
		if isthere == q.result[i] {
			return true
		}
		}  
		return false
	}

func main() {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.contains(2))
	fmt.Println()
	fmt.Println("Queue:", queue)

	element, ok := queue.Dequeue()
	if ok {
		fmt.Println("Dequeued:", element)
	}

	fmt.Println("Queue after dequeue:", queue)
	fmt.Println("Is queue empty?", queue.IsEmpty())
}
