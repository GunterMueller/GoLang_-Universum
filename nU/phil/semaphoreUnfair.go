package phil

// (c) Christian Maurer   v. 171229 - license see nU.go

// >>> Unfair solution with semaphores, danger of starvation

import
   "sync"
type
   semaphoreUnfair struct {
                    plate []sync.Mutex
                          }
var
  m sync.Mutex

func (x *semaphoreUnfair) test (p uint) {
  if status[p] == hungry &&
     status[left(p)] != dining && status[right(p)] != dining {
    changeStatus (p, dining)
    x.plate[p].Unlock()
  }
}

func newSU() Philos {
  x := new(semaphoreUnfair)
  x.plate = make([]sync.Mutex, 5)
  for p := uint(0); p < 5; p++ {
    x.plate[p].Lock()
  }
  return x
}

func (x *semaphoreUnfair) Lock (p uint) {
  changeStatus (p, hungry)
  m.Lock()
  x.test (p)
  m.Unlock()
  x.plate[p].Lock()
  changeStatus (p, dining)
}

func (x *semaphoreUnfair) Unlock (p uint) {
  changeStatus (p, thinking)
  m.Lock()
  x.test (left (p))
  x.test (right (p))
  m.Unlock()
}
