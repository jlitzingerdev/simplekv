package kvdb

type NodeStack struct {
	elements []*Node
}

func NewNodeStack() *NodeStack {
	return &NodeStack{}
}

func (stack *NodeStack) Push(node *Node) {
	stack.elements = append(stack.elements, node)
}

// Pop an element, panics if stack is empty
func (stack *NodeStack) Pop() *Node {
	el := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return el
}

func (stack *NodeStack) Empty() bool {
	return stack.elements == nil || len(stack.elements) == 0
}
