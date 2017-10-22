package sem

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> Implementation with a universal monitor

import (
  . "µU/obj"
  "µU/mon"
)
type
  monitor struct {
                 mon.Monitor
                 }

func newM (n uint) Semaphore {
  val := n
  x := new(monitor)
  f := func (a Any, i uint) Any {
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
