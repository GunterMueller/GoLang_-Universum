package rw

// (c) Christian Maurer   v. 241007 - license see µU.go

// >>> 1st readers/writers problem, solution with a baton

import
  "µU/sem"
type
  baton struct {
        nR, nW,     // number of active readers/writers
        bR, bW uint // number of blocked readers/writers
             e,     // the baton
          r, w sem.Semaphore
               }

func newB() ReaderWriter {
  x := new(baton)
  x.e = sem.New (1)
  x.r, x.w = sem.New (0), sem.New (0)
  return x
}

func (x *baton) vall() {
  if x.nW == 0 && x.bR > 0 {
    x.bR--
    x.r.V()
  } else if x.nR == 0 && x.nW == 0 && x.bW > 0 {
    x.bW--
    x.w.V()
  } else {
    x.e.V()
  }
}

func (x *baton) ReaderIn() {
  x.e.P()
  if x.nW > 0 {
    x.bR++
    x.e.V()
    x.r.P()
  }
  x.nR++
  x.vall()
}

func (x *baton) ReaderOut() {
  x.e.P()
  x.nR--
  x.vall()
}

func (x *baton) WriterIn() {
  x.e.P()
  if x.nR > 0 || x.nW > 0 {
    x.bW++
    x.e.V()
    x.w.P()
  }
  x.nW++
  x.vall()
}

func (x *baton) WriterOut() {
  x.e.P()
  x.nW--
  x.vall()
}

func (x *baton) Fin() {
}
