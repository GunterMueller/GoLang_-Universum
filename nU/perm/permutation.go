package perm

// (c) Christian Maurer   v. 171016 - license see nU.go

import
  "math/rand"
type
  permutation struct {
              uint "size"
            p []uint
              }

func rnd (n uint) uint {
  return uint(rand.Intn(int(n)))
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
    if rnd(rnd(1000)) % 2 == 1 {
      x.p[0], x.p[1] = x.p[1], x.p[0]
    }
  default:
    for i := uint(0); i < 3 * x.uint + rnd (x.uint); i++ {
      j, k := rnd (x.uint), rnd (x.uint)
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
