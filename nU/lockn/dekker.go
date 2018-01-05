package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

type dekker struct {
  interested [2]bool
  uint "Identität des begünstigten Prozesses, 0 oder 1"
}

func newDe() LockerN {
  return new(dekker)
}

func (x *dekker) Lock (p uint) { // p < 2
  if p > 1 { return }
  x.interested[p] = true
  for x.interested[1-p] {
    if x.uint == 1 - p {
      x.interested[p] = false
      for x.uint != p {
        nothing()
      }
      x.interested[p] = true
    }
  }
}

func (x *dekker) Unlock (p uint) { // p < 2
  if p > 1 { return }
  x.uint = 1 - p
  x.interested[p] = false
}
