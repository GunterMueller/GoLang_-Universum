package lockp

// (c) Christian Maurer   v. 161212 - license see murus.go

// >>> s. Ben-Ari S. 65

type
  doranthomas struct {
          interested [2]bool
            favoured uint // < 2
                     }

func newDT() LockerP {
  return new (doranthomas)
}

// Pre: p < 2
func (x *doranthomas) Lock (p uint) {
  if p > 1 { return }
  x.interested[p] = true
  if x.favoured == 1 - p {
    x.interested[p] = false
    for x.favoured != p { /* Null() */ } // await favoured == 1
    x.interested[p] = true
  }
  for x.interested[1-p] { /* Null() */ }
}

// Pre: p < 2
func (x *doranthomas) Unlock (p uint) {
  if p > 1 { return }
  x.interested[p] = false
  x.favoured = 1 - p
}
