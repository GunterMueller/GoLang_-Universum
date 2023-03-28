package symstk

// (c) Christian Maurer   v. 230303 - license see µU.go

import (
  "lsys/symstk/pair"
  "µU/stk"
)
var
  stack = stk.New(pair.New())

func push (b Symbol, i uint) {
  p := pair.New()
  p.Set (b, i)
  stack.Push (p)
}

func empty() bool {
  return stack.Empty()
}

func pop() (Symbol, uint) {
  if stack.Empty() {
    return '!', 0
  }
  p := stack.Pop().(pair.Pair)
  return p.Get()
}
