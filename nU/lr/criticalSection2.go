package lr

// (c) Christian Maurer   v. 171126 - license see nU.go

import "nU/cs"

type criticalSection2 struct {
  cs.CriticalSection
}

func newCS2() LeftRight {
  x := new(criticalSection2)
  var nL, nR uint
  c := func (i uint) bool {
         if i == left {
           return nR == 0 && (! x.Blocked (right) || nL == 0)
         }
         return nL == 0 && (! x.Blocked (left) || nR == 0)
       }
  e := func (i uint) uint {
         if i == left {
           nL++
           return nL
         } else {
           nR++
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

func (x *criticalSection2) LeftIn() {
  x.Enter (left)
}

func (x *criticalSection2) LeftOut() {
  x.Leave (left)
}

func (x *criticalSection2) RightIn() {
  x.Enter (right)
}

func (x *criticalSection2) RightOut() {
  x.Leave (right)
}
