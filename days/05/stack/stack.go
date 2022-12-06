package stack

type Stack interface {
	Push(value string)
	Pop() (value string, ok bool)
	Top() (value string, ok bool)
	Empty() bool
	Len() int
}

type StringStack struct {
	values []string
}

func NewStringStack(vs ...string) StringStack {
	return StringStack{
		values: vs,
	}
}

func (s *StringStack) Push(value string) {
	s.values = append(s.values, value)
}

func (s *StringStack) Pop() (value string, ok bool) {
	if s.Empty() {
		return "", false
	}

	value, ok = s.values[len(s.values)-1], true

	// remove the value
	s.values = s.values[:len(s.values)-1]

	return
}

func (s *StringStack) Top() (value string, ok bool) {
	if s.Empty() {
		return "", false
	}

	value, ok = s.values[len(s.values)-1], true
	return
}

func (s *StringStack) Empty() bool {
	return s.Len() < 1
}

func (s *StringStack) Len() int {
	return len(s.values)
}
