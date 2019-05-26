package col

// (c) Christian Maurer  v. 190402 - license see µU.go

import (
  . "µU/obj"
  "µU/ego"
)
type
  colour struct {
           level,     // of king
              id uint // identity of king"
                }
var
  c0 = C0()

func new_() Colour {
  x := new(colour)
  x.level, x.id = 0, ego.Me()
  return x
}

func (x *colour) imp (Y Any) *colour {
  y := Y.(*colour)
  return y
}

func (x *colour) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.level == y.level && x.id == y.id
}

func (x *colour) Copy (Y Any) {
  y := x.imp(Y)
  x.level, x.id = y.level, y.id
}

func (x *colour) Clone() Any {
  y := new_()
  y.Set (x.level, x.id)
  return y
}

func (x *colour) Less (Y Any) bool {
  y := x.imp(Y)
  if x.level == y.level {
    return x.id < y.id
  }
  return x.level < y.level
}

func (x *colour) Set (l, i uint) {
  x.level, x.id = l, i
}

func (x *colour) Level() uint {
  return x.level
}

func (x *colour) Id() uint {
  return x.id
}

func (x *colour) Inc() {
  x.level++
}

func (x *colour) Codelen() uint {
  return 2 * c0
}

func (x *colour) Encode() Stream {
  return append (Encode(x.level), Encode(x.id)...)
}

func (x *colour) Decode (s Stream) {
  x.level, x.id = Decode (0, s[:c0]).(uint), Decode (0, s[c0:]).(uint)
}
