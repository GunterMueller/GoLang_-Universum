package barr

// (c) Christian Maurer   v. 170628 - license see nU.go

import
  . "sync"
type
  gobarrier struct {
              uint "Anzahl der beteiligten Prozesse"
           waiting uint
                   *Cond
                   Mutex "to block"
                   }

func newG (m uint) Barrier {
  if m < 2 { return nil }
  x := new(gobarrier)
  x.uint = m
  x.Cond = NewCond(&x.Mutex)
  return x
}

func (x *gobarrier) Wait() {
  x.Mutex.Lock()
  x.waiting++
  if x.waiting < x.uint {
    x.Cond.Wait()
  } else {
    for x.waiting > 0 {
      x.waiting--
      if x.waiting > 0 {
        x.Cond.Signal()
      }
    }
  }
  x.Mutex.Unlock()
}
