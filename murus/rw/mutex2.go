package rw

// (c) murus.org  v. 170411 - license see murus.go

// >>> 2nd readers/writers problem (Courtois, Heymans, Parnas)
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 77

import
  . "sync"
type
  mutex2 struct {
         nR, bW int
 mutexR, mutexW,
      r, r1, rw Mutex
                }

func new2() ReaderWriter {
  return new (mutex2)
}

func (x *mutex2) ReaderIn() {
  x.r1.Lock()
  x.r.Lock()
  x.mutexR.Lock()
  x.nR++
  if x.nR == 1 {
    x.rw.Lock()
  }
  x.mutexR.Unlock()
  x.r.Unlock()
  x.r1.Unlock()
}

func (x *mutex2) ReaderOut() {
  x.mutexR.Lock()
  x.nR--
  if x.nR == 0 {
    x.rw.Unlock()
  }
  x.mutexR.Unlock()
}

func (x *mutex2) WriterIn() {
  x.mutexW.Lock()
  x.bW++
  if x.bW == 1 {
    x.r.Lock()
  }
  x.mutexW.Unlock()
  x.rw.Lock()
}

func (x *mutex2) WriterOut() {
  x.rw.Unlock()
  x.mutexW.Lock()
  x.bW--
  if x.bW == 0 {
    x.r.Unlock()
  }
  x.mutexW.Unlock()
}

func (x *mutex2) Fin() {
}
