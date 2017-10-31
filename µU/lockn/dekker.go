package lockn

// (c) Christian Maurer   v. 171024 - license see µU.go

import
  . "µU/obj"
type
  dekker struct {
     interested [2]bool
                uint "identity of the favoured process, < 2"
           inCS [2]bool
                }

func newDe() LockerN {
  return new(dekker)
}

func (x *dekker) chk (i uint) {
  for j := uint(0); j < 2; j++ {
    if j != i && x.inCS[j] {
      print("*")
      return
    }
  }
}

func (x *dekker) Lock (i uint) {
  x.interested[i] = true
  for x.interested[1-i] {
    if x.uint == 1 - i {
      x.interested[i] = false
      for x.uint != i {
        Gothing()
      }
      x.interested[i] = true
    }
  }
  x.inCS[i] = true; x.chk (i)
}

func (x *dekker) Unlock (i uint) {
  x.inCS[i] = false
  x.uint = 1 - i
  x.interested[i] = false
}
