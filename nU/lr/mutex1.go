package lr

// (c) Christian Maurer   v. 171125 - license see nU.go

import
  . "sync"
type
  mutex1 struct {
         nL, nR uint
     mL, mR, lr Mutex
                }

func new1() LeftRight {
  return new(mutex1)
}

func (x *mutex1) LeftIn() {
  x.mL.Lock()
  x.nL++
  if x.nL == 1 {
    x.lr.Lock()
  }
  x.mL.Unlock()
}

func (x *mutex1) LeftOut() {
  x.mL.Lock()
  x.nL--
  if x.nL == 0 {
    x.lr.Unlock()
  }
  x.mL.Unlock()
}

func (x *mutex1) RightIn() {
  x.mR.Lock()
  x.nR++
  if x.nR == 1 {
    x.lr.Lock()
  }
  x.mR.Unlock()
}

func (x *mutex1) RightOut() {
  x.mR.Lock()
  x.nR--
  if x.nR == 0 {
    x.lr.Unlock()
  }
  x.mR.Unlock()
}
