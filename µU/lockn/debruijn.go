package lockn

// (c) Christian Maurer   v. 171025 - license see µU.go

// >>> DeBruijn, J. G.: Additional Comments on a Problem in Concurrent Programming Control.
//     CACM 10 (1967) 137-138

import
  "runtime"
type
  deBruijn struct {
                  uint "number of processes"
         favoured uint
             flag []uint
                  }

func newDB (n uint) LockerN {
  x := new(deBruijn)
  x.uint = n
  x.flag = make([]uint, x.uint)
  return x
}

func (x *deBruijn) test (i uint) bool {
  for j := uint(0); j < x.uint; j++ {
    if j != i {
      if x.flag[j] == active { return false }
    }
  }
  return true
}

func (x *deBruijn) Lock (i uint) {
  for {
    x.flag[i] = requesting
    j := x.favoured
    for j != i {
      if x.flag[j] != passive {
        j = x.favoured
      } else {
        j = (j + x.uint - 1) % x.uint
      }
      runtime.Gosched()
    }
    x.flag[i] = active
    if x.test (i) { break }
  }
}

func (x *deBruijn) Unlock (i uint) {
  if x.flag[x.favoured] == passive || x.favoured == i {
    x.favoured = (x.favoured + x.uint - 1) % x.uint
  }
  x.flag[i] = passive
}
