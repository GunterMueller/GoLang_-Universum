package lock2

// (c) Christian Maurer   v. 190323 - license see nU.go

import (
  . "nU/atomic"
  . "nU/obj"
)
type
  peterson struct {
       interested [2]uint
         favoured uint "identity of the favoured process"
                  }

func newPeterson() Locker2 {
  return new(peterson)
}

func (x *peterson) Lock (p uint) {
  Store (&x.interested[p], 1)
  Store (&x.favoured, 1 - p)
  for x.interested[1-p] == 1 && x.favoured == 1-p {
    Nothing()
  }
}

func (x *peterson) Unlock (p uint) {
  Store (&x.interested[p], 0)
}
