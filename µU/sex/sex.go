package sex

// (c) Christian Maurer   v. 221003 - license see µU.go

import (
  . "µU/obj"
//  "µU/col"
  "µU/enum"
)
type
  sex struct {
             enum.Enum
             }

func new_() Sex {
  x := new (sex)
  x.Enum = enum.New (1)
  x.Enum.Set ("m", "d", "w")
  return x
}

func (x *sex) imp (Y any) *sex {
  y, ok := Y.(*sex)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *sex) Copy (Y any) {
  y := x.imp (Y)
  x.Enum.Copy (y.Enum)
}

func (x *sex) Eq (Y any) bool {
  return x.Enum.Eq (x.imp(Y).Enum)
}

func (x *sex) Less (Y any) bool {
  return false
}
