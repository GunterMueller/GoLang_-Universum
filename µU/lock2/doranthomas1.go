package lock2

// (c) Christian Maurer   v. 190815 - license see µU.go

import (
  . "µU/obj"
  . "µU/atomic"
)
type
  doranthomas1 struct {
                      uint "identity of the favoured process < 2"
           interested,
             afterYou [2]uint
             favoured uint
                      }

func newDoranThomas1() Locker2 {
  return new(doranthomas1)
}

func (x *doranthomas1) Lock (p uint) {
  Store (&x.interested[p], 1)
  if x.interested[1-p] == 1 {
    Store (&x.afterYou[p], 1)
    for x.interested[1-p] == 1 &&
        (x.favoured != p || x.afterYou[1-p] == 0) {
      Nothing()
    }
    Store (&x.afterYou[p], 0)
  }
}

func (x *doranthomas1) Unlock (p uint) {
  Store (&x.interested[p], 0)
  Store (&x.uint, 1-p)
}
