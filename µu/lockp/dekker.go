package lockp

// (c) Christian Maurer   v. 161212 - license see µu.go

type
  dekker struct {
     interested [2]bool
       favoured uint // < 2
                }

func newDe() LockerP {
  return new(dekker)
}

func (x *dekker) Lock (p uint) { // p < 2
  if p > 1 { return }
  x.interested[p] = true
  for x.interested[1-p] {
    if x.favoured == 1 - p {
      x.interested[p] = false
      for x.favoured != p { /* Null */ }
      x.interested[p] = true
    }
  }
}

func (x *dekker) Unlock (p uint) { // p < 2
  if p > 1 { return }
  x.favoured = 1 - p
  x.interested[p] = false
}

