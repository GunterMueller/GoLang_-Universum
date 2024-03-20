package audio

// (c) Christian Maurer   v. 240319 - license see µU.go

import
  . "µU/obj"
type
  Audio interface {

  Stringer
  TeXer
  Rotator

// Pre: y is of type Audio.
// Returns true, iff x is a part of y.
  Sub (Y any) bool
}

func New() Audio { return new_() }
