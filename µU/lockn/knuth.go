package lockn

// (c) Christian Maurer   v. 171025 - license see µU.go

// >>> Knuth, D. E.: Additional Comments on a Problem in Concurrent Programming Control.
//     CACM 9 (1966) 321-322

import
  "time"
const (
  passive = iota
  requesting
  active
)
type
  knuth struct {
               uint "number of processes"
      favoured uint
          flag []uint
               }

func newK (n uint) LockerN {
  x := new(knuth)
  x.uint = n
  x.favoured = x.uint
  x.flag = make([]uint, x.uint + 1)
  return x
}

func (x *knuth) test (i uint) bool {
  for j := uint(0); j < x.uint; j++ {
    if j != i {
      if x.flag[j] == active { return false }
    }
  }
  return true
}

func (x *knuth) Lock (i uint) {
  for {
    x.flag[i] = requesting
    j := x.favoured
    for j != i {
      if x.flag[j] == passive {
        j = (j + x.uint - 1) % x.uint
      } else {
        j = x.favoured
      }
      time.Sleep(1) // runtime.Gosched()
    }
    x.flag[i] = active
    if x.test (i) { break }
  }
  x.favoured = i
}

func (x *knuth) Unlock (i uint) {
  x.favoured = (i + x.uint - 1) % x.uint
  x.flag[i] = passive
}
