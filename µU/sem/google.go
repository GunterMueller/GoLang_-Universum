package sem

//     Christian Maurer   v. 171017

// (c) Go-Authors

import
  "sync"
type
  google struct {
            int "value"
          mutex *sync.Mutex
         wakeup *sync.Cond
                }

func newG (n uint) Semaphore {
  x := new(google)
  x.int = int(n)
  x.mutex = new(sync.Mutex)
  x.wakeup = sync.NewCond (x.mutex)
  return x
}

func (x *google) P() {
  x.mutex.Lock()
  x.int--
  for x.int < 0 {
    x.wakeup.Wait()
  }
  x.mutex.Unlock()
}

func (x *google) V() {
  x.mutex.Lock()
  wakeOthers := x.int < 0
  x.int++
  if wakeOthers {
    x.wakeup.Signal()
  }
  x.mutex.Unlock()
}
