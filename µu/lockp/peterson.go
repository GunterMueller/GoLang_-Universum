package lockp

// (c) Christian Maurer   v. 161212 - license see µu.go

type
  peterson struct {
       interested [2]bool
         favoured uint
                  }

func newP() LockerP {
  return new (peterson)
}

func (L *peterson) Lock (p uint) {
  if p > 1 { return }
  L.interested[p] = true
  L.favoured = 1-p
  for L.interested[1-p] && L.favoured == 1-p { /* do nothing */ }
}

func (L *peterson) Unlock (p uint) {
  if p > 1 { return }
  L.interested[p] = false
}
