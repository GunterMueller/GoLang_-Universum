package perm

// (c) Christian Maurer   v. 161216 - license see µu.go

import
  "µu/rand"
type
  permutation struct {
              uint "size"
            p []uint
              }

func new_(n uint) Permutation {
  if n == 0 { return nil }
  x := new (permutation)
  x.uint = n
  if x.uint > 1 {
    x.p = make ([]uint, x.uint)
    for i := uint(0); i < x.uint; i++ {
      x.p[i] = i
    }
  }
  x.Permute()
  return x
}

func (x *permutation) Permute() {
  switch x.uint {
  case 1:
    return
  case 2:
    if rand.Natural (rand.Natural (1000)) % 2 == 1 {
      x.p[0], x.p[1] = x.p[1], x.p[0]
    }
  default:
    for i := uint(0); i < 3 * x.uint + rand.Natural (x.uint); i++ {
      j, k := rand.Natural (x.uint), rand.Natural (x.uint)
      if j != k {
        x.p[j], x.p[k] = x.p[k], x.p[j]
      }
    }
  }
}

func (x *permutation) F (i uint) uint {
  if x.uint == 1 || x.uint <= i {
    return 0
  }
  return x.p[i]
}
