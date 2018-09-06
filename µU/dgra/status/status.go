package status

// (c) Christian Maurer  v. 180902 - license see µU.go

import (
  . "µU/obj"
  "µU/ego"
)
type
  status struct {
                int "phase of king"
                uint "identity of king"
                }

func new_() Status {
  x := new(status)
  x.int, x.uint = -1, ego.Me()
  return x
}

func (x *status) imp (Y Any) *status {
  y := Y.(*status)
  return y
}

func (x *status) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.int == y.int && x.uint == y.uint
}

func (x *status) Copy (Y Any) {
  y := x.imp(Y)
  x.int, x.uint = y.int, y.uint
}

func (x *status) Clone() Any {
  y := new_()
  y.Set (x.int, x.uint)
  return y
}

func (x *status) Less (Y Any) bool {
  y := x.imp(Y)
  if x.int == y.int {
    return x.uint < y.uint
  }
  return x.int < y.int
}

func (x *status) Set (p int, i uint) {
  x.int, x.uint = p, i
}

func (x *status) Phase() int {
  return x.int
}

func (x *status) Id() uint {
  return x.uint
}

func (x *status) Inc() {
  x.int++
}

func (x *status) Codelen() uint {
  return 2 * C0()
}

func (x *status) Encode() Stream {
  return append (Encode(x.int), Encode(x.uint)...)
}

func (x *status) Decode (s Stream) {
  x.int, x.uint = Decode (0, s[:C0()]).(int), Decode (0, s[C0():]).(uint)
}
