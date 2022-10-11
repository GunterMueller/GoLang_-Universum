package mbbuf

// (c) Christian Maurer   v. 220702 - license see nU.go

import ("nU/bbuf"; "nU/mon")

type monitor struct {
  mon.Monitor
}

func newM (a any, n uint) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  buffer := bbuf.New (a, n)
  x := new(monitor)
  f := func (a any, i uint) any {
         if i == ins {
           buffer.Ins (a)
           x.Monitor.Signal (ins)
           return nil
         }
         if buffer.Empty() {
           x.Monitor.Wait (ins)
         }
         return buffer.Get()
       }
  x.Monitor = mon.New (nFuncs, f)
  return x
}

func (x *monitor) Ins (a any) {
  x.Monitor.F (a, ins)
}

func (x *monitor) Get() any {
  return x.Monitor.F (nil, get)
}
