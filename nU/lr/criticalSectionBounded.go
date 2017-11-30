package lr

// (c) Christian Maurer   v. 171126 - license see nU.go

import "nU/cs"

type criticalSectionBounded struct {
  cs.CriticalSection
}

func newCSB (mL, mR uint) LeftRight {
  x := new(criticalSectionBounded)
  if mL == 0 { mL = 1 }
  if mR == 0 { mR = 1 }
  var nL, nR,
      tL, tR uint // number of lefties/righties within one turn
  c := func (i uint) bool {
         if i == left {
           return nR == 0 && (! x.Blocked (right) || tL < mL)
         }
         return nL == 0 && (! x.Blocked (left) || tR < mR)
       }
  e := func (i uint) uint {
         if i == left {
           nL++
           tL++
           tR = 0
           return nL
         } else {
           nR++
           tR++
           tL = 0
         }
         return nR
       }
  l := func (i uint) {
         if i == left {
           nL--
         } else {
           nR--
         }
       }
  x.CriticalSection = cs.New (2, c, e, l)
  return x
}

func (x *criticalSectionBounded) LeftIn() {
  x.Enter (left)
}

func (x *criticalSectionBounded) LeftOut() {
  x.Leave (left)
}

func (x *criticalSectionBounded) RightIn() {
  x.Enter (right)
}

func (x *criticalSectionBounded) RightOut() {
  x.Leave (right)
}
