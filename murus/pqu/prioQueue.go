package pqu

// (c) murus.org  v. 150122 - license see murus.go

import (
  . "murus/obj"
  "murus/pqu/internal"
)
const
  pack = "pqu"
type
  prioQueue struct {
            object Any
            anchor internal.Heap
               num uint
                   }

func New (a Any) PrioQueue {
  x:= new (prioQueue)
  x.object = Clone (a)
  x.anchor = internal.New()
  return x
}

func (x *prioQueue) Num() uint {
  return x.num
}

func (x *prioQueue) Ins (a Any) {
  CheckTypeEq (a, x.object)
  x.num ++
  x.anchor = x.anchor.Ins (a, x.num).(internal.Heap)
  x.anchor.Lift (x.num)
}

func (x *prioQueue) Get() Any {
  if x.num == 0 {
    return nil
  }
  return x.anchor.Get()
}

func (x *prioQueue) Del() Any {
  if x.num == 0 {
    return nil
  }
  if x.num == 1 {
    a:= x.anchor.Get()
    x.anchor = internal.New()
    x.num = 0
    return a
  }
  y, a:= x.anchor.Del (x.num)
  x.anchor = y.(internal.Heap)
  x.num --
  if x.num > 0 {
    x.anchor.Sift (x.num)
  }
  return a
}
