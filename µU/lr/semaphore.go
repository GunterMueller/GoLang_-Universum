package lr

// (c) Christian Maurer   v. 171015 - license see µU.go

// >>> 1st left/right problem

import
  "µU/sem"
type
  semaphore struct {
            nL, nR uint
           m, l, r sem.Semaphore
                   }

func newS() LeftRight {
  x := new(semaphore)
  x.m = sem.New (1)
  x.l = sem.New (1)
  x.r = sem.New (1)
  return x
}

func (x *semaphore) LeftIn() {
  x.l.P()
  x.nL++
  if x.nL == 1 {
    x.m.P()
  }
  x.l.V()
}

func (x *semaphore) LeftOut() {
  x.l.P()
  x.nL--
  if x.nL == 0 {
    x.m.V()
  }
  x.l.V()
}

func (x *semaphore) RightIn() {
  x.r.P()
  x.nR++
  if x.nR == 1 {
    x.m.P()
  }
  x.r.V()
}

func (x *semaphore) RightOut() {
  x.r.P()
  x.nR--
  if x.nR == 0 {
    x.m.V()
  }
  x.r.V()
}

func (x *semaphore) Fin() {
}
