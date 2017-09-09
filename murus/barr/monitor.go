package barr

// (c) Christian Maurer   v. 140330 - license see murus.go

// >>> barrierlementation with a monitor
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 162

import (
  . "murus/obj"
  "murus/mon"
)
type
  barrierMon struct {
           involved,
            waiting uint
                    mon.Monitor
                    }

func newMon (n uint) Barrier {
  x := new (barrierMon)
  if n < 2 { return nil }
  x.involved = n
  f := func (a Any, i uint) Any {
         x.waiting++
         if x.waiting < x.involved {
           x.Monitor.Wait(0)
         } else {
/*
           x.Monitor.SignalAll(0) // simple solution with broadcast
           x.waiting = 0
*/
           for x.waiting > 0 { // standard solution
             x.waiting --
             if x.waiting > 0 {
               x.Monitor.Signal(0)
             }
           }
         }
         return nil
       }
  x.Monitor = mon.New (1, f, nil)
  return x
}

func (x *barrierMon) Wait() {
  x.F (nil, 0)
}
