package stack

// (c) Christian Maurer   v. 230225 - license see µU.go

import (
  "µU/vect"
  "µU/stk"
)
var (
  stack = stk.New (vect.New())
  stack1 = stk.New (vect.New())
)

func empty() bool {
  return stack.Empty()
}

func push (v vect.Vector) {
  stack.Push (v)
}

func pop() vect.Vector {
  if stack.Empty() { return vect.New() }
  return stack.Pop().(vect.Vector)
}

func empty1() bool {
  return stack1.Empty()
}

func push1 (v vect.Vector) {
  stack1.Push (v)
}

func pop1() vect.Vector {
  if stack1.Empty() { return vect.New() }
  return stack1.Pop().(vect.Vector)
}
