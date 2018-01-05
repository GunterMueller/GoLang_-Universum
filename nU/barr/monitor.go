package barr

// (c) Christian Maurer   v. 171019 - license see nU.go

import (. "nU/obj"; "nU/mon")

type monitor struct {
  mon.Monitor
}

func newM (n uint) Barrier {
  involved := n
  waiting := uint(0)
  x := new(monitor)
  if n < 2 { return nil }
  f := func (a Any, i uint) Any {
         waiting++
         if waiting < involved {
           x.Monitor.Wait(0)
         } else {
           for waiting > 0 { // Standard-Lösung
             waiting--
             if waiting > 0 {
               x.Monitor.Signal(0)
             }
           }
         }
         return waiting
       }
  x.Monitor = mon.New (1, f)
  return x
}

func (x *monitor) Wait() {
  x.F (nil, 0)
}
