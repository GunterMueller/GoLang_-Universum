package fahrweg

// (c) Christian Maurer   v. 230109 - license see µU.go

import (
  . "bahn/richtung"
  "bahn/block"
)
type
  weg struct {
           n []uint
             }

func new_() Weg {
  x := new(weg)
  x.n = make([]uint, 0)
  return x
}

func (x *weg) Start() uint {
  return x.n[0]
}

func (x *weg) Ziel() uint {
  return x.n[len(x.n)-1]
}

func (x *weg) Clr() {
  x.n = make([]uint, 0)
}

func (x *weg) Insert (n uint) {
  x.n = append (x.n, n)
}

func (x *weg) Nr (i uint) uint {
  return x.n[i]
}

func (x *weg) Num() uint {
  return uint(len(x.n))
}

func (x *weg) Less (i, j int) bool {
  return x.n[i] < x.n[j]
}

func (x *weg) Leq (i, j int) bool {
  return x.n[i] <= x.n[j]
}

func (x *weg) Ablenkend() bool {
  for i := 1; i < len(x.n); i++ {
    b := block.B[i]
    if b.IstWeiche() || b.IstDKW() {
      if b.Stellung() != Gerade {
        return true
      }
    }
  }
  return false
}
