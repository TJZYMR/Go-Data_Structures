package main

import "fmt"

type Stack struct {
	stack []int
	Top   int //index of the top element
}

//empty() – Returns whether the stack is empty – Time Complexity: O(1)
// size() – Returns the size of the stack – Time Complexity: O(1)
// top() – Returns a reference to the topmost element of the stack – Time Complexity: O(1)
// push(a) – Inserts the element ‘a’ at the top of the stack – Time Complexity: O(1)
// pop() – Deletes the topmost element of the stack – Time Complexity: O(1)

func (S *Stack) Push(a int) {
	if len(S.stack) == 0 {
		S.stack = append(S.stack, a)
		S.Top = 0
	} else {
		S.stack = append(S.stack, a)
		S.Top = S.Top + 1
	}

}
func (S *Stack) Pop() {
	if len(S.stack) == 0 {
		fmt.Println("Stack is empty cannot pop")
	} else if len(S.stack) == 1 {
		S.stack = []int{}
		S.Top = -1
	} else {
		S.stack = S.stack[:len(S.stack)-1]
		S.Top = len(S.stack) - 1
	}
}

//implemeting stack data structure
func main() {
	s := []int{} //initially empty stack
	stack := Stack{stack: s, Top: 0}
	stack.Push(11)
	stack.Push(22)
	stack.Push(77)
	fmt.Println("Stack Status:==>", stack)

	stack.Pop() //giving index of the element to be deleted
	//stack.Pop(2) //giving index of the element to be deleted
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Push(11)
	stack.Push(22)
	stack.Push(77)
	fmt.Println("Stack Status:==>", stack)
	fmt.Print("Top Element:=>")
	if stack.Top == -1 {
		fmt.Println("-1,", "Stack is empty")
	} else {
		fmt.Println("Top Element", stack.stack[stack.Top], "With index:=>", stack.Top)
	}

}
