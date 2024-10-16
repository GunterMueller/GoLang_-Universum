package sem

// (c) Christian Maurer   v. 241001 - license see µU.go

// >>> corrected solution with more liveliness

import
  "sync"
type
  semaphore1 struct {
                int "value"
              block,
              mutex,
                seq sync.Mutex
                  n int
                    }

func new1 (n uint) Semaphore {
  x := new(semaphore1)
  x.int = int(n)
  x.block.Lock()
  return x
}

func (x *semaphore1) P() {
  x.seq.Lock()
  x.mutex.Lock()
  x.int--
  if x.int < 0 {
    x.mutex.Unlock()
    x.block.Lock()
    x.mutex.Lock()
    x.n--
    if x.n > 0 {
      x.block.Unlock()
    }
  }
  x.mutex.Unlock()
  x.seq.Unlock()
}

func (x *semaphore1) V() {
  x.mutex.Lock()
  x.int++
  if x.int <= 0 {
    x.block.Unlock()
  } else {
    x.mutex.Unlock()
  }
}
