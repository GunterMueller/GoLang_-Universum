package smok

// (c) murus.org  v. 170629 - license see murus.go

// >>> Naive solution with deadlock danger

import (
  "sync"
  . "murus/smok/utensil"
)
type
  naive struct {
     smokerOut sync.Mutex
         avail [NUtensils]sync.Mutex
               }

func new_() Smokers {
  x := new (naive)
  x.smokerOut.Lock()
  for u := uint(0); u < NUtensils; u++ {
    x.avail[u].Lock()
  }
  return x
}

func (x *naive) Agent (u uint) {
  x.smokerOut.Lock()
  u1, u2 := Others(u)
  x.avail[u1].Unlock()
  x.avail[u2].Unlock()
}

func (x *naive) SmokerIn (u uint) {
  u1, u2 := Others(u)
  x.avail[u1].Lock()
  x.avail[u2].Lock()
}

func (x *naive) SmokerOut() {
  x.smokerOut.Unlock()
}
