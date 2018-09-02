package status

// (c) Christian Maurer  v. 180819 - license see µU.go

import
  . "nU/obj"
type
  status struct {
      phase, id uint // of king
                }

func new_(p, i uint) Status {
  x := new(status)
  x.phase, x.id = p, i
  return x
}

func (x *status) imp (Y Any) *status {
  y := Y.(*status)
  return y
}

func (x *status) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.phase  == y.phase && x.id == y.id
}

func (x *status) Copy (Y Any) {
  y := x.imp(Y)
  x.phase, x.id = y.phase, y.id
}

func (x *status) Clone() Any {
  return new_(x.phase, x.id)
}

func (x *status) Less (Y Any) bool {
  y := x.imp(Y)
  if x.phase  == y.phase {
    return x.id < y.id
  }
  return x.phase < y.phase
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

func (x *status) Codelen() uint {
  return 2 * C0
}

func (x *status) Encode() Stream {
  return append (Encode(x.phase), Encode(x.id)...)
}

func (x *status) Decode (s Stream) {
  x.phase, x.id = Decode (0, s[:C0]).(uint), Decode (0, s[C0:]).(uint)
}
