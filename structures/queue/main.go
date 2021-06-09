package main

type Stack struct {
	items []int
}

func (s *Stack) Push(el int) {
	s.items = append(s.items, el)
}

func (s *Stack) Pop(el int) int {
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value
}

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(el int) {
	q.items = append(q.items, el)
}

func (q *Queue) Dequeue(el int) int {
	value := q.items[0]
	q.items = q.items[1:]
	return value
}

func main() {

}
