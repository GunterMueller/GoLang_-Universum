package lr

// (c) murus.org  v. 170731 - license see murus.go

// >>> left/right problem: implementation with critical sections
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 93

import (
  . "murus/obj"
  "murus/cs"
)
type
  criticalSectionBounded struct {
                         nL, nR, // number of active lefties/righties
                         tL, tR uint // number of lefties/righties within one turn
                                cs.CriticalSection
                                }

func newCSB (mL, mR uint) LeftRight {
  x := new (criticalSectionBounded)
  if mL == 0 { mL = 1 }
  if mR == 0 { mR = 1 }
  c := func (k uint) bool {
         switch k {
         case left:
           return x.nR == 0 && (! x.Blocked (right) || x.tL < mL)
         case right:
           return x.nL == 0 && (! x.Blocked (left) || x.tR < mR)
         }
         panic("")
       }
  i := func (A Any, k uint) {
         switch k {
         case left:
           x.nL++
           writeL (x.nL)
           x.tL++
           x.tR = 0
         case right:
           x.nR++
           writeR (x.nR)
           x.tR++
           x.tL = 0
         }
       }
  o := func (A Any, k uint) {
         switch k {
         case left:
           x.nL--
           writeL (x.nL)
         case right:
           x.nR--
           writeR (x.nR)
         }
       }
  x.CriticalSection = cs.New (2, c, i, o)
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

func (x *criticalSectionBounded) Fin() {
}
