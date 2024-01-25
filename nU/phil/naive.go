package phil

// (c) Christian Maurer   v. 171229 - license see nU.go

// >>> Naive solution with deadlock

import
  "sync"
type
  naive struct {
          fork []sync.Mutex
               }

func new_() Philos {
  x := new(naive)
  x.fork = make([]sync.Mutex, 5)
  return x
}

func (x *naive) Lock (p uint) {
  changeStatus (p, hungry)
  x.fork[left (p)].Lock()
  changeStatus (p, hasLeftFork)
  x.fork[p].Lock()
  changeStatus (p, dining)
}

func (x *naive) Unlock (p uint) {
  changeStatus (p, thinking)
  x.fork[p].Unlock()
  x.fork[left (p)].Unlock()
}
