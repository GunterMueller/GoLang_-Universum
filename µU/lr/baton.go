package lr

// (c) Christian Maurer   v. 241007 - license see µU.go

// >>> 1st left/right problem; implementation with a baton

import
  "µU/sem"
type
  baton struct {
        nL, nR,     // number of active lefties/righties
        bL, bR uint // number of blocked lefties/righties
             e,     // the baton
          l, r sem.Semaphore
               }

func newB() LeftRight {
  x := new(baton)
  x.e = sem.New (1)
  x.l, x.r = sem.New (0), sem.New (0)
  return x
}

func (x *baton) vall() {
  if x.nL == 0 && x.bR > 0 {
    x.bR--
    x.r.V()
  } else if x.nR == 0 && x.bL > 0 {
    x.bL--
    x.l.V()
  } else {
    x.e.V()
  }
}

func (x *baton) LeftIn() {
  x.e.P()
  if x.nR > 0 {
    x.bL++
    x.e.V()
    x.l.P()
  }
  x.nL++
  x.vall()
}

func (x *baton) LeftOut() {
  x.e.P()
  x.nL--
  x.vall()
}

func (x *baton) RightIn() {
  x.e.P()
  if x.nL > 0 {
    x.bR++
    x.e.V()
    x.r.P()
  }
  x.nR++
  x.vall()
}

func (x *baton) RightOut() {
  x.e.P()
  x.nR--
  x.vall()
}

func (x *baton) Fin() {
}
