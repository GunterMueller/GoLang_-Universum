package pqu

// (c) Christian Maurer   v. 170316 - license see murus.go

import (
  . "murus/obj"
  "murus/pqu/internal"
)
type
  prioQueue struct {
                   Any "to maintain the type of objects in the queue"
                   internal.Heap "classical structure"
                   uint "number of objects in the queue"
                   }

func new_(a Any) PrioQueue {
  x := new(prioQueue)
  x.Any = Clone(a)
  x.Heap = internal.New()
  return x
}

func (x *prioQueue) Empty() bool {
  return x.uint == 0
}

func (x *prioQueue) Num() uint {
  return x.uint
}

func (x *prioQueue) Ins (a Any) {
  CheckTypeEq (a, x.Any)
  x.uint++
  x.Heap = x.Heap.Ins (a, x.uint).(internal.Heap)
  x.Heap.Lift (x.uint)
}

func (x *prioQueue) Get() Any {
  if x.uint == 0 {
    return nil
  }
  return x.Heap.Get()
}

func (x *prioQueue) Del() Any {
  if x.uint == 0 {
    return nil
  }
  if x.uint == 1 {
    a := x.Heap.Get()
    x.Heap = internal.New()
    x.uint = 0
    return a
  }
  y, a := x.Heap.Del (x.uint)
  x.Heap = y.(internal.Heap)
  x.uint--
  if x.uint > 0 {
    x.Heap.Sift (x.uint)
  }
  return a
}
