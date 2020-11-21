package stack

// (c) Christian Maurer   v. 201103 - license see µU.go

import (
  "µU/ker"
  "µU/stk"
  "µU/spc/stack/internal"
)
var (
  stack = stk.New (nums.New(N))
  stack1 = stk.New (nums.New(N))
)

func empty() bool {
  return stack.Empty()
}

func push (r ...float64) {
  if len(r) != N { ker.Panic ("wrong number of float64's") }
  x := nums.New (N)
  x.Set (r...)
  stack.Push (x)
}

func pop() []float64 {
  s := nums.New (N)
  if ! stack.Empty() {
    s = stack.Pop().(nums.Numbers)
  }
  return s.Get()
}

func empty1() bool {
  return stack1.Empty()
}

func push1 (r ...float64) {
  if len(r) != N { ker.Panic ("wrong number of float64's") }
  x := nums.New (N)
  x.Set (r...)
  stack1.Push (x)
}

func pop1() []float64 {
  s := nums.New (N)
  if ! stack1.Empty() {
    s = stack1.Pop().(nums.Numbers)
  }
  return s.Get()
}
