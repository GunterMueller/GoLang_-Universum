package lockn

// (c) Christian Maurer   v. 171024 - license see µU.go

// >>> s. Ben-Ari S. 65

import (
  . "µU/obj"
)
type
  doranthomas struct {
          interested [2]bool
                     uint "identity of the favoured process < 2"
                     }

func newDT() LockerN {
  return new(doranthomas)
}

// Pre: p < 2
func (x *doranthomas) Lock (p uint) {
  x.interested[p] = true
  if x.uint == 1 - p {
    x.interested[p] = false
    for x.uint != p { // await uint == 1
      Gothing()
    }
    x.interested[p] = true
  }
  for x.interested[1-p] {
    Gothing()
  }
}

// Pre: p < 2
func (x *doranthomas) Unlock (p uint) {
  x.interested[p] = false
  x.uint = 1 - p
}
