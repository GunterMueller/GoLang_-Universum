package cr

// (c) Christian Maurer   v. 170423 - license see µu.go

import (
  "µu/ker"
  . "µu/obj"
  "µu/cs"
)
type (
  status struct {
            max []uint // indexed over process classes
         number,
          class uint
                }
  criticalResource struct {
                     stat []status // indexed over resources
                   nC, nR uint
                       cs.CriticalSection
                          }
)

func new_(nc, nr uint) CriticalResource {
  x := new (criticalResource)
  x.nC, x.nR = nc, nr
  x.stat = make ([]status, x.nC)
  for r := uint(0); r < x.nR; r++ {
    x.stat[r].max = make ([]uint, x.nC)
    for c := uint(0); c < x.nC; c++ {
      x.stat[r].max[c] = ker.MaxNat()
    }
  }
  c := func (k uint) bool {
        var b bool
        for r := uint(0); r < x.nR; r++ {
          b = b ||
              x.stat[r].number == 0 ||
              x.stat[r].class == k && x.stat[r].number < x.stat[r].max[k]
        }
        return b
      }
/*
  p := func (a Any, k uint) bool { // see cs experiment
        var b bool
        for r := uint(0); r < x.nR; r++ {
          b = b ||
              x.stat[r].number == 0 ||
              x.stat[r].class == k && x.stat[r].number < x.stat[r].max[k]
        }
        return b
      }
*/
  e := func (a Any, k uint) {
        for r := uint(0); r < x.nR; r++ {
          if x.stat[r].number == 0 || x.stat[r].class == k {
            x.stat[r].class = k
            x.stat[r].number ++
            n := a.(*uint)
            *n = r
            return
          }
        }
        ker.Oops()
      }
  l := func (a Any, k uint) {
        for r := uint(0); r < x.nR; r++ {
          if x.stat[r].class == k && x.stat[r].number > 0 {
            x.stat[r].number --
          }
        }
      }
  x.CriticalSection = cs.New (x.nC, c, e, l)
  return x
}

func (x *criticalResource) Limit (m [][]uint) {
  for c := uint(0); c < x.nC; c++ {
    for r := uint(0); r < x.nR; r++ {
      x.stat[r].max[c] = m[c][r]
    }
  }
}

func (x *criticalResource) Enter (k uint) uint {
  var r uint
  x.CriticalSection.Enter (k, &r)
  return r
}

func (x *criticalResource) Leave (k uint) {
  x.CriticalSection.Leave (k, 0)
}
