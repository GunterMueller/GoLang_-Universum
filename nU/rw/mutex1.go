package rw

// (c) Christian Maurer   v. 170411 - license see nU.go

import
  . "sync"
type
  mutex1 struct {
             nR uint
          m, rw Mutex
                }

func new1() ReaderWriter {
  return new(mutex1)
}

func (x *mutex1) ReaderIn() {
  x.m.Lock()
  x.nR++
  if x.nR == 1 {
    x.rw.Lock()
  }
  x.m.Unlock()
}

func (x *mutex1) ReaderOut() {
  x.m.Lock()
  x.nR--
  if x.nR == 0 {
    x.rw.Unlock()
  }
  x.m.Unlock()
}

func (x *mutex1) WriterIn() {
  x.rw.Lock()
}

func (x *mutex1) WriterOut() {
  x.rw.Unlock()
}
