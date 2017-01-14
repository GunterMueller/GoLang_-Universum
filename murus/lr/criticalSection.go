package lr

// (c) murus.org  v. 150304 - license see murus.go

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

func NewCriticalSection() LeftRight {
  x:= new (criticalSection)
  c:= func (i uint) bool { return x.n[1 - i] == 0 }
  e:= func (a Any, i uint) { x.n[i]++ }
  l:= func (a Any, i uint) { x.n[i]-- }
  x.CriticalSection = cs.New (2, c, e, l)
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
