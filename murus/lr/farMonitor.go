package lr

// (c) murus.org  v. 161225 - license see murus.go

// >>> left/right problem: implementation with far monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 209

import (
  . "murus/obj"
  "murus/host"
  "murus/fmon"
)
type
  farMonitor struct {
             nL, nR uint
                    fmon.FarMonitor
                    }

func NewFarMonitor (h host.Host, p uint16, s bool) LeftRight {
  x := new(farMonitor)
  c := func (a Any, i uint) bool {
        switch i {
        case lIn:
          return x.nR == 0
        case rIn:
          return x.nL == 0
        }
        return true // lOut, rOut
      }
  f := func (a Any, i uint) Any {
        switch i {
        case lIn:
          x.nL++
        case rIn:
          x.nR++
        case lOut:
          x.nL--
        case rOut:
          x.nR--
        }
        return true
      }
  x.FarMonitor = fmon.New (true, nFuncs, f, c, h, p, s)
  return x
}

func (x *farMonitor) LeftIn() {
  x.F (true, lIn)
}

func (x *farMonitor) LeftOut() {
  x.F (true, lOut)
}

func (x *farMonitor) RightIn() {
  x.F (true, rIn)
}

func (x *farMonitor) RightOut() {
  x.F (true, rOut)
}
