package mbbuf

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "µU/bbuf"
  "µU/mon"
)
type
  mbufM struct {
               mon.Monitor
               }

func newM (a any, n uint) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  buffer := bbuf.New (a, n)
  x := new(mbufM)
  f := func (a any, i uint) any {
         if i == ins {
           buffer.Ins (a)
           x.Monitor.Signal (ins)
           return nil
         }
         if buffer.Num() == 0 {
           x.Monitor.Wait (ins)
         }
         return buffer.Get()
       }
  x.Monitor = mon.New (nFuncs, f)
  return x
}

func (x *mbufM) Ins (a any) {
  x.Monitor.F (a, ins)
}

func (x *mbufM) Get() any {
  return x.Monitor.F (nil, get)
}
