package sem

// (c) murus.org  v. 140803 - license see murus.go

// >>> Algorithm of Barz
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 73

import
  "sync"
type
  barz struct {
    cs, mutex sync.Mutex
              int "value"
              }

func NewBarz (n uint) Semaphore {
  x:= new (barz)
  x.int = int(n)
  if x.int == 0 {
    x.cs.Lock()
  }
  return x
}

func (x *barz) P() {
  x.cs.Lock()
  x.mutex.Lock()
  x.int--
  if x.int > 0 {
    x.cs.Unlock()
  }
  x.mutex.Unlock()
}

func (x *barz) V() {
  x.mutex.Lock()
  x.int++
  if x.int == 1 {
    x.cs.Unlock()
  }
  x.mutex.Unlock()
}
