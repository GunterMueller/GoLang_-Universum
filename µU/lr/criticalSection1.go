package lr

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> 1st left/right problem

import
  "µU/cs"
type
  criticalSection1 struct {
                          cs.CriticalSection
                          }

func newCS1() LeftRight {
  var nL, nR uint
  x := new(criticalSection1)
  c := func (i uint) bool {
         if i == left {
           return nR == 0
         }
         return nL == 0
       }
  e := func (i uint) uint {
         if i == left {
           nL++
           return nL
         }
         nR++
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

func (x *criticalSection1) LeftIn() {
  _ = x.Enter (left)
}

func (x *criticalSection1) LeftOut() {
  x.Leave (left)
}

func (x *criticalSection1) RightIn() {
  _ = x.Enter (right)
}

func (x *criticalSection1) RightOut() {
  x.Leave (right)
}

func (x *criticalSection1) Fin() {
}
