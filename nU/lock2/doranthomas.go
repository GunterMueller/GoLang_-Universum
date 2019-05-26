package lock2

// (c) Christian Maurer   v. 190323 - license see nU.go

// >>> s. Ben-Ari S. 65

import (
  . "nU/atomic"
  . "nU/obj"
)
type
  doranthomas struct {
          interested [2]uint
                     uint "identity of the favoured process < 2"
                     }

func newDoranThomas() Locker2 {
  return new(doranthomas)
}

// Pre: p < 2
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

// Pre: p < 2
func (x *doranthomas) Unlock (p uint) {
  Store (&x.interested[p], 0)
  Store (&x.uint, 1-p)
}
