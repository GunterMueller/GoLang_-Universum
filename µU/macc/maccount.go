package macc

// (c) Christian Maurer   v. 170121 - license see µU.go

import (
  "sync"
  "µU/euro"
)
type
  mAccount struct {
                  euro.Euro "balance"
                  sync.Mutex
                  *sync.Cond
                  }

func new_() MAccount {
  x:= new (mAccount)
  x.Euro = euro.New()
  x.Euro.Set2 (0, 0)
  x.Cond = sync.NewCond (&x.Mutex)
  return x
}

func (x *mAccount) Deposit (e euro.Euro) euro.Euro {
  x.Mutex.Lock()
  x.Euro.Add (e)
  x.Cond.Signal()
  x.Mutex.Unlock()
  return x.Euro.Clone().(euro.Euro)
}

func (x *mAccount) Draw (e euro.Euro) euro.Euro {
  x.Mutex.Lock()
  for x.Euro.Less(e) {
    x.Cond.Wait()
  }
  x.Euro.Sub(e)
  x.Cond.Signal()
  x.Mutex.Unlock()
  return x.Euro.Clone().(euro.Euro)
}
