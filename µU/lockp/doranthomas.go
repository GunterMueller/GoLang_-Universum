package lockp

// (c) Christian Maurer   v. 171013 - license see µU.go

// >>> s. Ben-Ari S. 65

type
  doranthomas struct {
          interested [2]bool
                     uint "identity of the favoured process < 2"
                     }

func newDT() LockerP {
  return new(doranthomas)
}

// Pre: p < 2
func (x *doranthomas) Lock (p uint) {
  if p > 1 { return }
  x.interested[p] = true
  if x.uint == 1 - p {
    x.interested[p] = false
    for x.uint != p { /* Null() */ } // await uint == 1
    x.interested[p] = true
  }
  for x.interested[1-p] { /* Null() */ }
}

// Pre: p < 2
func (x *doranthomas) Unlock (p uint) {
  if p > 1 { return }
  x.interested[p] = false
  x.uint = 1 - p
}
