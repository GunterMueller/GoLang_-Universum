package mstk

// (c) murus.org  v. 161216 - license see murus.go

import (
  "sync"
  . "murus/obj"
  "murus/stk"
)
type
  mStack struct {
                stk.Stack
       notEmpty,
          mutex sync.Mutex
                }

func newMstk (a Any) MStack {
  if a == nil { return nil } // XXX
  x:= new(mStack)
  x.Stack = stk.New(a)
  x.notEmpty.Lock()
  return x
}

func (x *mStack) Push (a Any) {
  x.mutex.Lock()
  x.Stack.Push (a)
  x.mutex.Unlock()
  x.notEmpty.Unlock()
}

func (x *mStack) Pop() {
  x.notEmpty.Lock()
  x.mutex.Lock()
  x.Stack.Pop()
  x.mutex.Unlock()
}

func (x *mStack) Top() Any {
  x.notEmpty.Lock()
  x.mutex.Lock()
  defer x.mutex.Unlock()
  return x.Stack.Top()
}
