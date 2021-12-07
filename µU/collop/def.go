package collop

// (c) Christian Maurer   v. 211126 - license see µU.go

import
  . "µU/obj"

// Pre: f returns true, iff x is a part of y.
func Operate (c Collector, o Indexer, f func (x, y Indexer) bool, t *string) { operate (c,o,f,t) }
