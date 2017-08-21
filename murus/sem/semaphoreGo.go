package sem

//     murus.org  v. 140803
//
// (c) Go-Authors

import
  "sync"
type
  semaphoreGo struct {
                     int "value"
               mutex *sync.Mutex
              wakeup *sync.Cond
                     }

func newGo (n int) Semaphore {
  x := new(semaphoreGo)
  x.int = n
  x.mutex = new(sync.Mutex)
  x.wakeup = sync.NewCond (x.mutex)
  return x
}

func (x *semaphoreGo) P() {
  x.mutex.Lock()
  x.int--
  for x.int < 0 {
    x.wakeup.Wait()
  }
  x.mutex.Unlock()
}

func (x *semaphoreGo) V() {
  x.mutex.Lock()
  wakeOthers := x.int < 0
  x.int++
  if wakeOthers {
    x.wakeup.Signal()
  }
  x.mutex.Unlock()
}
