package lockp

// (c) Christian Maurer   v. 161212 - license see µU.go

type
  peterson struct {
       interested [2]bool
                  uint "identity of the favoured process"
                  }

func newP() LockerP {
  return new(peterson)
}

func (x *peterson) Lock (p uint) {
  if p > 1 { return }
  x.interested[p] = true
  x.uint = 1 - p
  for x.interested[1-p] && x.uint == 1 - p { /* Null() */ }
}

func (x *peterson) Unlock (p uint) {
  if p > 1 { return }
  x.interested[p] = false
}
