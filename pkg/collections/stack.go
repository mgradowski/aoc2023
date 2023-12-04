package collections

import "errors"

var ErrStackEmpty = errors.New("cannot pop from an empty stack")

type Stack[T any] struct {
	buf []T
	len int
}

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

func (stack *Stack[T]) Push(els ...T) {
	for _, el := range els {
		if len(stack.buf) == stack.len {
			stack.buf = append(stack.buf, el)
		} else {
			stack.buf[stack.len] = el
		}
		stack.len++
	}
}

func (stack *Stack[T]) Pop() (T, error) {
	if stack.len == 0 {
		var zero T
		return zero, ErrStackEmpty
	}
	el := stack.buf[stack.len-1]
	stack.len--
	return el, nil
}

func (stack *Stack[T]) Len() int {
	return stack.len
}
