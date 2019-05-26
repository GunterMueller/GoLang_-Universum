package lock2

// (c) Christian Maurer   v. 190331 - license see µU.go

// >>> s. Ben-Ari S. 65

import (
  . "µU/atomic"
  . "µU/obj"
)
type
  doranthomas struct {
                     uint "identity of the favoured process < 2"
          interested [2]uint
            favoured uint
                     }

func newDoranThomas() Locker2 {
  return new(doranthomas)
}

func (x *doranthomas) Lock (p uint) {
  Store (&x.interested[p], 1)
  if x.uint == 1-p {
    Store (&x.interested[p], 0)
    for x.uint != p {
      Nothing()
    }
    Store (&x.interested[p], 1)
  }
  for x.interested[1-p] == 1 {
    Nothing()
  }
}

func (x *doranthomas) Unlock (p uint) {
  Store (&x.interested[p], 0)
  Store (&x.uint, 1-p)
}
