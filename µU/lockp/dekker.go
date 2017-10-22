package lockp

// (c) Christian Maurer   v. 171013 - license see µU.go

type
  dekker struct {
     interested [2]bool
                uint "identity of the favoured process, < 2"
                }

func newDe() LockerP {
  return new(dekker)
}

func (x *dekker) Lock (p uint) { // p < 2
  if p > 1 { return }
  x.interested[p] = true
  for x.interested[1-p] {
    if x.uint == 1 - p {
      x.interested[p] = false
      for x.uint != p { /* Null */ }
      x.interested[p] = true
    }
  }
}

func (x *dekker) Unlock (p uint) { // p < 2
  if p > 1 { return }
  x.uint = 1 - p
  x.interested[p] = false
}

