package lr

// (c) Christian Maurer   v. 170731 - license see murus.go

// >>> Left/Right problem: Simple Solution with mutexes
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 79

import
  . "sync"
type
  leftRight struct {
            nL, nR int
            mL, mR,
                lr Mutex
                   }

func new_() LeftRight {
  return new (leftRight)
}

func (x *leftRight) LeftIn() {
  x.mL.Lock()
  x.nL++
  writeL (x.nL)
  if x.nL == 1 {
    x.lr.Lock()
  }
  x.mL.Unlock()
}

func (x *leftRight) LeftOut() {
  x.mL.Lock()
  x.nL--
  writeL (x.nL)
  if x.nL == 0 {
    x.lr.Unlock()
  }
  x.mL.Unlock()
}

func (x *leftRight) RightIn() {
  x.mR.Lock()
  x.nR++
  writeR (x.nR)
  if x.nR == 1 {
    x.lr.Lock()
  }
  x.mR.Unlock()
}

func (x *leftRight) RightOut() {
  x.mR.Lock()
  x.nR--
  writeR (x.nR)
  if x.nR == 0 {
    x.lr.Unlock()
  }
  x.mR.Unlock()
}

func (x *leftRight) Fin() {
}
