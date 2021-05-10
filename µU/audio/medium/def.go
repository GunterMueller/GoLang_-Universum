package medium

// (c) Christian Maurer   v. 210509 - license see µU.go

import
  . "µU/obj"

type
  Medium interface {
    Object
    Editor
    TeX() string
  }

func New() Medium { return new_() }
