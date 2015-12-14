package queue

import "fmt"

type Queue struct {
	data []int
}

func Push(q *Queue, n int) {
	q.data = append(q.data, n)
}

func Front(q *Queue) (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.data[0], nil
}

func Pop(q *Queue) (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return_value := q.data[0]
	q.data = q.data[1:]
	return return_value, nil
}

func IsEmpty(q *Queue) bool {
	return len(q.data) == 0
}
