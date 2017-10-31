package acc

// (c) Christian Maurer   v. 171020 - license see µU.go

import
  "sync"
type
  account struct {
                 uint "balance"
                 sync.Mutex
                 *sync.Cond
                 }

func new_() Account {
  x := new(account)
  x.Cond = sync.NewCond (&x.Mutex)
  return x
}

func (x *account) Deposit (a uint) uint {
  x.Mutex.Lock(); defer x.Mutex.Unlock()
  x.uint += a
  x.Cond.Signal()
  return x.uint
}

func (x *account) Draw (a uint) uint {
  x.Mutex.Lock(); defer x.Mutex.Unlock()
  for x.uint < a {
    x.Cond.Wait()
  }
  x.uint -= a
  x.Cond.Signal()
  return x.uint
}
