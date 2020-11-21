package mstk

// (c) Christian Maurer   v. 201103 - license see µU.go

import (
  "sync"
  . "µU/obj"
  "µU/stk"
)
type
  mStack struct {
                stk.Stack
       notEmpty,
          mutex sync.Mutex
                }

func new_(a Any) MStack {
  CheckAtomicOrObject(a)
  x := new(mStack)
  x.Stack = stk.New(a)
  x.notEmpty.Lock()
  return x
}

func (x *mStack) Empty() bool {
  x.mutex.Lock()
  defer x.mutex.Unlock()
  return x.Stack.Empty()
}

func (x *mStack) Push (a Any) {
  x.mutex.Lock()
  x.Stack.Push (a)
  x.mutex.Unlock()
  x.notEmpty.Unlock()
}

func (x *mStack) Pop() Any {
  x.notEmpty.Lock()
  x.mutex.Lock()
  defer x.mutex.Unlock()
  return x.Stack.Pop()
}
