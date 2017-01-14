package lr

// (c) murus.org  v. 150914 - license see murus.go

// >>> left/right problem: implementation with critical sections
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 93

import (
  "murus/ker"
  . "murus/obj"
  "murus/cs"
)
type
  criticalSectionBounded struct {
                      n, nr, nz [2]uint
                                cs.CriticalSection
                                }

func NewCriticalSectionBounded (l, r uint) LeftRight {
  x:= new (criticalSectionBounded)
  if l == 0 { l = ker.MaxNat() }
  if r == 0 { r = ker.MaxNat() }
  x.nz[0], x.nz[1] = l, r
  c:= func (i uint) bool {
        return x.n[1-i] == 0 &&
               (! x.Blocked (1-i) || x.nr[i] < x.nz[i])
      }
  e:= func (A Any, i uint) {
        x.n[i]++
        x.nr[i]++
        x.nr[1-i] = 0
      }
  a:= func (A Any, i uint) {
        x.n[i]--
      }
  x.CriticalSection = cs.New (2, c, e, a)
  return x
}

func (x *criticalSectionBounded) LeftIn() {
  x.Enter (left, nil)
}

func (x *criticalSectionBounded) LeftOut() {
  x.Leave (left, nil)
}

func (x *criticalSectionBounded) RightIn() {
  x.Enter (right, nil)
}

func (x *criticalSectionBounded) RightOut() {
  x.Leave (right, nil)
}
