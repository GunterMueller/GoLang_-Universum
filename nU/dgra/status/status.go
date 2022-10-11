package status

// (c) Christian Maurer  v. 220801 - license see nU.go

import (
  . "nU/obj"
  "nU/ego"
)

type status struct {
  phase, id uint // of king
}

func new_() Status {
  x := new(status)
  x.phase, x.id = 0, ego.Me()
  return x
}

func (x *status) imp (Y any) *status {
  y := Y.(*status)
  return y
}

func (x *status) Eq (Y any) bool {
  y := x.imp(Y)
  return x.phase  == y.phase && x.id == y.id
}

func (x *status) Copy (Y any) {
  y := x.imp(Y)
  x.phase, x.id = y.phase, y.id
}

func (x *status) Clone() any {
  y := new_()
  y.Set (x.phase, x.id)
  return y
}

func (x *status) Less (Y any) bool {
  y := x.imp(Y)
  if x.phase  == y.phase {
    return x.id < y.id
  }
  return x.phase < y.phase
}

func (x *status) Set (p, i uint) {
  x.phase, x.id = p, i
}

func (x *status) Phase() uint {
  return x.phase
}

func (x *status) Id() uint {
  return x.id
}

func (x *status) Inc() {
  x.phase++
}

var c0 = C0()

func (x *status) Codelen() uint {
  return 2 * c0
}

func (x *status) Encode() Stream {
  return append (Encode(x.phase), Encode(x.id)...)
}

func (x *status) Decode (s Stream) {
  x.phase, x.id = Decode (0, s[:c0]).(uint), Decode (0, s[c0:]).(uint)
}
