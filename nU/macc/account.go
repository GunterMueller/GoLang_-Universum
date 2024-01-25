package macc

// (c) Christian Maurer   v. 171020 - license see nU.go

import
   "sync"
type
  maccount struct {
                  uint "balance"
                  sync.Mutex
                  *sync.Cond
                  }

func new_() MAccount {
  x := new(maccount)
  x.Cond = sync.NewCond (&x.Mutex)
  return x
}

func (x *maccount) Deposit (a uint) uint {
  x.Mutex.Lock(); defer x.Mutex.Unlock()
  x.uint += a
  x.Cond.Signal()
  return x.uint
}

func (x *maccount) Draw (a uint) uint {
  x.Mutex.Lock(); defer x.Mutex.Unlock()
  for x.uint < a {
    x.Cond.Wait()
  }
  x.uint -= a
  x.Cond.Signal()
  return x.uint
}
