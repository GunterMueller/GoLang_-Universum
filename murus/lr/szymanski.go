package lr

// (c) Christian Maurer   v. 170731 - license see murus.go

// >>> Left/Right problem: Simple Solution with Szymanskis Lock.

import (
  "murus/lockp"
  "murus/ego"
)
type
  lrSzy struct {
        nL, nR int
        mL, mR,
            lr lockp.LockerP
               }
var
  me = ego.Me()

func newSz() LeftRight {
  x := new(lrSzy)
  x.mL = lockp.NewSzymanski(4)
  x.mR = lockp.NewSzymanski(4)
  x.lr = lockp.NewSzymanski(4)
  return x
}

func (x *lrSzy) LeftIn() {
  x.mL.Lock(me)
  x.nL++
  writeL (x.nL)
  if x.nL == 1 {
    x.lr.Lock(me)
  }
  x.mL.Unlock(me)
}

func (x *lrSzy) LeftOut() {
  x.mL.Lock(me)
  x.nL--
  writeL (x.nL)
  if x.nL == 0 {
    x.lr.Unlock(me)
  }
  x.mL.Unlock(me)
}

func (x *lrSzy) RightIn() {
  x.mR.Lock(me)
  x.nR++
  writeR (x.nR)
  if x.nR == 1 {
    x.lr.Lock(me)
  }
  x.mR.Unlock(me)
}

func (x *lrSzy) RightOut() {
  x.mR.Lock(me)
  x.nR--
  writeR (x.nR)
  if x.nR == 0 {
    x.lr.Unlock(me)
  }
  x.mR.Unlock(me)
}

func (x *lrSzy) Fin() {
}
