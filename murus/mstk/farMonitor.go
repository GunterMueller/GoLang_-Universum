package mstk

// (c) murus.org  v. 161226 - license see murus.go

import (
  . "murus/obj"
  "murus/stk"
  "murus/host"
  "murus/fmon"
)
const (
  push = uint(iota)
  pop
  top
  nFuncs
)
type
  farMonitor struct {
                    stk.Stack
                    fmon.FarMonitor
                    }

func NewFarMonitor (a Any, h host.Host, p uint16, s bool) MStack {
  x := new (farMonitor)
  x.Stack = stk.New (a)
  c := func (a Any, i uint) bool {
         if i == push {
           return true
         }
         return ! x.Stack.Empty() // top, pop
       }
  f := func (a Any, i uint) Any {
         switch i {
         case push:
           x.Stack.Push (a)
         case pop:
           x.Stack.Pop()
         case top:
           return x.Stack.Top()
         }
         return a // push, pop
       }
  x.FarMonitor = fmon.New (a, nFuncs, f, c, h, p, s)
  return x
}

func (x *farMonitor) Push (a Any) {
  x.FarMonitor.F (a, push)
}

func (x *farMonitor) Pop() {
  var a Any
  x.FarMonitor.F(a, pop)
}

func (x *farMonitor) Top() Any {
  var a Any
  return x.FarMonitor.F(a, top)
}
