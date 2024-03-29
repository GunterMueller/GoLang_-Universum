package collop

// (c) Christian Maurer   v. 211214 - license see µU.go

import
  . "µU/obj"

// Pre: f returns true, iff x is a part of y.
func Operate (c Collector, o Rotator, f func (x, y Rotator) bool) { operate (c,o,f) }
