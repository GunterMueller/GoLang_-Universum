package rw

// (c) murus.org  v. 170411 - license see murus.go

// >>> 1st readers/writers problem (readers' preference)
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 75

import
  . "sync"
type
  mutex1 struct {
                nR int
      mutex, rw Mutex
                }

func new1() ReaderWriter {
  return new (mutex1)
}

func (x *mutex1) ReaderIn() {
  x.mutex.Lock()
  x.nR++
  if x.nR == 1 {
    x.rw.Lock()
  }
  x.mutex.Unlock()
}

func (x *mutex1) ReaderOut() {
  x.mutex.Lock()
  x.nR--
  if x.nR == 0 {
    x.rw.Unlock()
  }
  x.mutex.Unlock()
}

func (x *mutex1) WriterIn() {
  x.rw.Lock()
}

func (x *mutex1) WriterOut() {
  x.rw.Unlock()
}

func (x *mutex1) Fin() {
}
