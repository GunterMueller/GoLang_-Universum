package lr

// (c) Christian Maurer   v. 170731 - license see murus.go

// >>> left/right problem: implementation with critical sections
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 85 ff.

import (
  . "murus/obj"
  "murus/cs"
)
type
  criticalSection struct {
                       n [2]uint
                         cs.CriticalSection
                         }

func newCS() LeftRight {
  x := new (criticalSection)
  c := func (i uint) bool { return x.n[1 - i] == 0 }
  es := func (a Any, i uint) {
          x.n[i]++
          if i == left { writeL (x.n[i]) } else { writeR (x.n[i]) }
        }
  ls := func (a Any, i uint) {
          x.n[i]--
          if i == left { writeL (x.n[i]) } else { writeR (x.n[i]) }
        }
  x.CriticalSection = cs.New (2, c, es, ls)
  return x
}

func (x *criticalSection) LeftIn() {
  x.Enter (left, nil)
}

func (x *criticalSection) LeftOut() {
  x.Leave (left, nil)
}

func (x *criticalSection) RightIn() {
  x.Enter (right, nil)
}

func (x *criticalSection) RightOut() {
  x.Leave (right, nil)
}

func (x *criticalSection) Fin() {
}
