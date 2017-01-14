package mon

// (c) murus.org  v. 150303 - license see murus.go

import (
  . "murus/ker"; . "murus/obj"
//  "murus/errh"
)
const
  pack = "mon"
type
  farMonitor struct {
                    Any
                    uint "number of monitor functions"
                 ch []chan Any
                    FuncSpectrum
                    PredSpectrum
                    Stmt
                    }

func NewF (a Any, n uint, f FuncSpectrum, p PredSpectrum) Monitor {
  if n == 0 { Panic ("mon.NewF must be called with 2nd arg > 0") }
  x:= new (farMonitor)
  x.Any = Clone (a)
  x.uint = n
  x.ch = make ([]chan Any, x.uint)
  x.FuncSpectrum, x.PredSpectrum, x.Stmt = f, p, Null
  for i:= uint(0); i < x.uint; i++ {
    x.ch[i] = make (chan Any)
  }
  x.Stmt()
  for i:= uint(0); i < x.uint; i++ {
    go func (j uint) {
      select {
      case a, p:= <-When (x.PredSpectrum (x.Any, j), x.ch[j]):
        if p {
          x.ch[j] <- x.FuncSpectrum (a, j)
        }
      }
    }(i)
  }
  return x
}

/*
func (x *farMonitor) Prepare (s Stmt) {
  x.Stmt = s
}
*/

func (x *farMonitor) Wait (i uint) {
// TODO
}

func (x *farMonitor) Awaited (i uint) bool {
  return false // TODO
}

func (x *farMonitor) Signal (i uint) {
// TODO
}

func (x *farMonitor) SignalAll (i uint) {
// TODO
}

func (x *farMonitor) F (a Any, i uint) Any {
  x.ch[i] <- a
  return <-x.ch[i]
}

func (x *farMonitor) S (a Any, i uint, c chan Any) {
// TODO
}
