package lockn

// (c) Christian Maurer   v. 171024 - license see µU.go

import
  . "µU/obj"
type
  peterson struct {
       interested [2]bool
                  uint "identity of the favoured process"
                  }

func newP() LockerN {
  return new(peterson)
}

func (x *peterson) Lock (i uint) {
  x.interested[i] = true
  x.uint = 1 - i
  for x.interested[1-i] && x.uint == 1 - i {
    Gothing()
  }
}

func (x *peterson) Unlock (i uint) {
  x.interested[i] = false
}
