package lockn

// (c) Christian Maurer   v. 171024 - license see µU.go

//     T H I S   A L G O R I T H M   I S   N O T   C O R R E C T

import
  . "µU/obj"
type
  kessels2 struct {
       interested [2]bool
         favoured [2]uint // < 2
                  }

func newK2() LockerN {
  return new(kessels2)
}

func (x *kessels2) Lock (i uint) {
  x.interested[i] = true
  local := (i + x.favoured[1-i]) % 2
  x.favoured[i] = local
  for x.interested[1-i] && local == (i + x.favoured[1-i]) % 2 {
    Gothing()
  }
}

func (x *kessels2) Unlock (i uint) {
  x.interested[i] = false
}
