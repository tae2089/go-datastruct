package stack

import "errors"

type Stack struct {
	data []int
}

func CreateStack() *Stack {
	data := make([]int, 0)
	stack := &Stack{
		data: data,
	}
	return stack

}

func (stack *Stack) Push(value int) {
	stack.data = append(stack.data, value)
}

func (stack *Stack) Pop() (int, error) {
	dataLength := len(stack.data)
	if dataLength == 0 {
		return -1, errors.New("No Data")
	}
	item, data := stack.data[dataLength-1], stack.data[0:dataLength-1]
	stack.data = data
	return item, nil

}
