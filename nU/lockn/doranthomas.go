
package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

type doranthomas struct {
  interested [2]bool
  uint "Identität des begünstigten Prozesses, 0 oder 1"
}

func newDT() LockerN {
  return new(doranthomas)
}

// Vor.: p < 2
func (x *doranthomas) Lock (p uint) {
  if p > 1 { return }
  x.interested[p] = true
  if x.uint == 1 - p {
    x.interested[p] = false
    for x.uint != p {
      nothing()
    }
    x.interested[p] = true
  }
  for x.interested[1-p] {
    nothing()
  }
}

// Vor.: p < 2
func (x *doranthomas) Unlock (p uint) {
  if p > 1 { return }
  x.interested[p] = false
  x.uint = 1 - p
}
