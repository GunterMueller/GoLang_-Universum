package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

type kessels2 struct {
  interested [2]bool
  favoured [2]uint // < 2
}

func newK2() LockerN {
  return new(kessels2)
}

func (x *kessels2) Lock (p uint) {
  if p > 1 { return }
  x.interested[p] = true
  local:= (p + x.favoured[1 - p]) % 2
  x.favoured[p] = local
  for x.interested[1 - p] && local == (p + x.favoured[1 - p]) % 2 {
    nothing()
  }
}

func (x *kessels2) Unlock (p uint) {
  if p > 1 { return }
  x.interested[p] = false
}
