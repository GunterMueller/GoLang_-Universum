package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

type peterson struct {
  interested [2]bool
  uint "Identität des begünstigten Prozesses, 0 oder 1"
}

func newP() LockerN {
  return new(peterson)
}

func (x *peterson) Lock (p uint) {
  if p > 1 { return }
  x.interested[p] = true
  x.uint = 1 - p
  for x.interested[1-p] && x.uint == 1 - p {
    nothing()
  }
}

func (x *peterson) Unlock (p uint) {
  if p > 1 { return }
  x.interested[p] = false
}
