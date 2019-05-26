package lock2

// (c) Christian Maurer   v. 190323 - license see µU.go

import (
  . "µU/atomic"
  . "µU/obj"
)
type
  dekker struct {
     interested [2]uint
                uint "identity of the favoured process, < 2"
                }

func newDekker() Locker2 {
  return new(dekker)
}

func (x *dekker) Lock (p uint) {
  Store (&x.interested[p], 1)
  for x.interested[1-p] == 1 {
    if x.uint == 1 - p {
      Store (&x.interested[p], 0)
      for x.uint != p {
        Nothing()
      }
      Store (&x.interested[p], 1)
    }
    Nothing()
  }
}

func (x *dekker) Unlock (p uint) {
  Store (&x.uint, 1 - p)
  Store (&x.interested[p], 0)
}
