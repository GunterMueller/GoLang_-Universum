package lr

// (c) murus.org  v. 140330 - license see murus.go

// >>> Left/Right problem: Simple Solution with leftRightes
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 79

import
  . "sync"
type
  leftRight struct {
            nL, nR int
            mL, mR,
                lr Mutex
                   }

func New() LeftRight {
  return new (leftRight)
}

func (x *leftRight) LeftIn() {
  x.mL.Lock()
  x.nL++
  if x.nL == 1 {
    x.lr.Lock()
  }
  x.mL.Unlock()
}

func (x *leftRight) LeftOut() {
  x.mL.Lock()
  x.nL--
  if x.nL == 0 {
    x.lr.Unlock()
  }
  x.mL.Unlock()
}

func (x *leftRight) RightIn() {
  x.mR.Lock()
  x.nR++
  if x.nR == 1 {
    x.lr.Lock()
  }
  x.mR.Unlock()
}

func (x *leftRight) RightOut() {
  x.mR.Lock()
  x.nR--
  if x.nR == 0 {
    x.lr.Unlock()
  }
  x.mR.Unlock()
}
