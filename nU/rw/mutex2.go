package rw

// (c) Christian Maurer   v. 171125 - license see nU.go

import
  . "sync"
type
  mutex2 struct {
         nR, bW uint
         mR, mW,
      r, r1, rw Mutex
                }

func new2() ReaderWriter {
  return new(mutex2)
}

func (x *mutex2) ReaderIn() {
  x.r1.Lock()
  x.r.Lock()
  x.mR.Lock()
  x.nR++
  if x.nR == 1 {
    x.rw.Lock()
  }
  x.mR.Unlock()
  x.r.Unlock()
  x.r1.Unlock()
}

func (x *mutex2) ReaderOut() {
  x.mR.Lock()
  x.nR--
  if x.nR == 0 {
    x.rw.Unlock()
  }
  x.mR.Unlock()
}

func (x *mutex2) WriterIn() {
  x.mW.Lock()
  x.bW++
  if x.bW == 1 {
    x.r.Lock()
  }
  x.mW.Unlock()
  x.rw.Lock()
}

func (x *mutex2) WriterOut() {
  x.rw.Unlock()
  x.mW.Lock()
  x.bW--
  if x.bW == 0 {
    x.r.Unlock()
  }
  x.mW.Unlock()
}
