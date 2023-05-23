package main 


import (
	"fmt" 
)

// Stack represents stack that hold a slice 
type Stack struct { 
	items []int 

}


// The `Push`  function  will add a value at the end . 
func (s *Stack) Push (i int) { 
	s.items = append(s.items, i)  
}



// The `Pop` function will remove  a value at the end and will 
// return the removed value . 
// we need a pointer receiver
func (s *Stack) Pop () int { 
	l := len(s.items) - 1 
	toRemove := s.items[l] 
	s.items = s.items[:l] 
	return toRemove 	
}

func main() { 
	myStack := Stack{}
	fmt.Println("The initial stack " , myStack) 
	myStack.Push(100)
	myStack.Push(200) 
	myStack.Push(300) 
	myStack.Push(400) 
	myStack.Push(500)
	myStack.Push(600)
	fmt.Println("The stack is now ->",  myStack)
	myStack.Pop()  
	fmt.Println("The stack is now ->",  myStack)
	myStack.Pop() 
	fmt.Println("The stack is now ->",  myStack)
	myStack.Pop()
	fmt.Println("The stack is now ->",  myStack)
	myStack.Push(400)
	fmt.Println("The stack is now ->",  myStack)
	

}