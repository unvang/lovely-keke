package stack

type stack struct {
	top    *node
	Length int
}
type node struct {
	val  int
	prev *node
}

func NewStack() *stack {
	return &stack{nil, 0}
}
func (s *stack) Top() int {
	return s.top.val
}
func (s *stack) Pop() {
	cur := s.top
	s.top = cur.prev
	s.Length--
}
func (s *stack) Push(val int) {
	n := &node{val, s.top}
	s.top = n
	s.Length++
}
