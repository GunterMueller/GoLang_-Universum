package barr

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> implementation with a monitor

import (
  . "µU/obj"
  "µU/mon"
)
type
  monitor struct {
                 mon.Monitor
                 }

func newMon (n uint) Barrier {
  involved, waiting := n, uint(0)
  x := new(monitor)
  if n < 2 { return nil }
  f := func (a Any, i uint) Any {
         waiting++
         if waiting < involved {
           x.Monitor.Wait(0)
         } else {
/*
           x.Monitor.SignalAll(0) // simple solution with broadcast
           x.waiting = 0
*/
           for waiting > 0 { // standard solution
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
