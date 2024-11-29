package mstk

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/stk"
  "µU/fmon"
)
const (
  push = uint(iota)
  pop
  nFuncs
)
type
  farMonitor struct {
                    stk.Stack
                    fmon.FarMonitor
                    }

func newFM (a any, h string, p uint16, s bool) MStack {
  CheckAtomicOrObject (a)
  x := new (farMonitor)
  x.Stack = stk.New (a)
  c := func (a any, i uint) bool {
         if i == push {
           return true
         }
         return ! x.Stack.Empty() // pop
       }
  f := func (a any, i uint) any {
         if i == push {
           x.Stack.Push (a)
           return a
         }
         return x.Stack.Pop()
       }
  x.FarMonitor = fmon.New (a, nFuncs, f, c, h, p, s)
  return x
}

func (x *farMonitor) Push (a any) {
  x.FarMonitor.F (a, push)
}

func (x *farMonitor) Pop() any {
  var a any
  return x.FarMonitor.F(a, pop)
}
