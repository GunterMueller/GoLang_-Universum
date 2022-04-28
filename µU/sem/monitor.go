package sem

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> Implementation with a universal monitor

import (
  "µU/mon"
)
type
  monitor struct {
                 mon.Monitor
                 }

func newM (n uint) Semaphore {
  val := n
  x := new(monitor)
  f := func (a any, i uint) any {
         if i == p {
           if val == 0 {
             x.Monitor.Wait (v)
           }
           val--
         } else {
           val++
           x.Monitor.Signal (v)
         }
         return val
       }
  x.Monitor = mon.New (2, f)
  return x
}

func (x *monitor) P() {
  x.F (nil, p)
}

func (x *monitor) V() {
  x.F (nil, v)
}
