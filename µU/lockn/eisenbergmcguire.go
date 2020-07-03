package lockn

// (c) Christian Maurer   v. 191229 - license see µU.go

// >>> Eisenberg, M. A., McGuire, M. R.: Further comments on Dijkstra's
//                       concurrent programming control problem.
//     CACM 15 (1966) 999

import (
  . "µU/obj"
  . "µU/atomic"
)
type
  eisenbergMcGuire struct {
                          uint "number of processes"
                 favoured uint
                     flag []uint
                          }

func (x *eisenbergMcGuire) Lock (p uint) {
  for {
    Store (&x.flag[p], requesting)
    j := x.favoured
    for j != p {
      if x.flag[j] == passive {
        j = (j + x.uint - 1) % x.uint
      } else {
        j = x.favoured
      }
    }
    Nothing()
    j = 0
    Store (&x.flag[p], active)
    for j < x.uint && (j == 0 || x.flag[j] != active) {
      j++
    }
    if j <= x.uint && (x.favoured == p || x.flag[x.favoured] == passive) {
      break
    }
  }
  Store (&x.favoured, p)
}

func (x *eisenbergMcGuire) Unlock (p uint) {
  j := (x.favoured + 1) % x.uint
  for j != x.favoured && (x.flag[j] == passive) {
    j = (j + 1) % x.uint
  }
  Store (&x.favoured, j)
  Store (&x.flag[p], passive)
}

func newEisenbergMcGuire (n uint) LockerN {
  x := new(knuth)
  x.uint = n
  x.favoured = x.uint
  x.flag = make([]uint, x.uint + 1)
  return x
}
