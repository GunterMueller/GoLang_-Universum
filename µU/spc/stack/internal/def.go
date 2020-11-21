package nums

// (c) Christian Maurer  v. 201101 - liecense see µU.go

import
  . "µU/obj"
type
  Numbers interface {

  Object

  Set (x ...float64)
  Get() []float64
}

func New (n uint) Numbers { return new_(n) }
