package smok

// (c) Christian Maurer   v. 170629 - license see µu.go

// >>> Solution with helper processes due to D. L. Parnas:
//     On a Solution to the Cigarette Smoker's Problem
//     Comm. ACM 18 (1975), 181-183

import (
  "sync"
  . "µu/smok/utensil"
)
type
  parnas struct {
          avail [NUtensils]bool
          mutex,
          agent sync.Mutex
       supplied,
          smoke [NUtensils]sync.Mutex
                }

func (x *parnas) help (u uint) {
  var first bool
  for {
    x.supplied[u].Lock()
    x.mutex.Lock()
    u1, u2 := Others(u)
    first = true
    if x.avail[u1] {
      first = false
      x.avail[u1] = false
      x.smoke[u2].Unlock()
    }
    if x.avail[u2] {
      first = false
      x.avail[u2] = false
      x.smoke[u1].Unlock() }
    if first {
      x.avail[u] = true
    }
    x.mutex.Unlock()
  }
}

func newP() Smokers {
  x := new (parnas)
  for u := uint(0); u < NUtensils; u++ {
    x.supplied[u].Lock()
    x.smoke[u].Lock()
  }
  x.agent.Lock()
  for u := uint(0); u < NUtensils; u++ {
    go x.help (u)
  }
  return x
}

func (x *parnas) Agent (u uint) {
  x.agent.Lock()
  u1, u2 := Others(u)
  x.supplied[u1].Unlock()
  x.supplied[u2].Unlock()
}

func (x *parnas) SmokerIn (u uint) {
  x.smoke[u].Lock()
}

func (x *parnas) SmokerOut() {
  x.agent.Unlock()
}
