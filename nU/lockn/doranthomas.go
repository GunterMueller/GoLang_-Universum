package lockn

// (c) Christian Maurer   v. 190316 - license see nU.go

// >>> s. Ben-Ari S. 65

import
  . "nU/obj"
type
  doranthomas struct {
          interested [2]bool
                     uint "identity of the favoured process < 2"
                     }

func newDoranThomas() LockerN {
  return new(doranthomas)
}

// Pre: p < 2
func (x *doranthomas) Lock (p uint) {
  x.interested[p] = true
  if x.uint == 1 - p {
    x.interested[p] = false
    for x.uint != p {
      Nothing()
    }
    x.interested[p] = true
  }
  for x.interested[1-p] {
    Nothing()
  }
}

// Pre: p < 2
func (x *doranthomas) Unlock (p uint) {
  x.interested[p] = false
  x.uint = 1 - p
}
