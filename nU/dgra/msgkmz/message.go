package msgkmz

// (c) Christian Maurer   v. 180819 - license see µU.go

import (
  . "nU/obj"
  "nU/dgra/status"
)
type
  message struct {
                 byte "message kind"
               s status.Status
                 }

func new_(k byte, s status.Status) Message {
  x := new(message)
  x.byte, x.s = k, s
  return x
}

func (x *message) imp (Y Any) *message {
  y, _ := Y.(*message)
  return y
}


func (x *message) Kind() byte {
  return x.byte
}

func (x *message) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.byte == y.byte && x.s.Eq (y.s)
}

func (x *message) Copy (Y Any) {
  y := x.imp(Y)
  x.byte = y.byte
  x.s.Copy (y.s)
}

func (x *message) Clone() Any {
  return new_(x.byte, x.s)
}

func (x *message) Status() status.Status {
  return x.s
}

func (x *message) Codelen() uint {
  return 1 + x.s.Codelen()
}

func (x *message) Encode() Stream {
  return append(Encode(x.byte), x.s.Encode()...)
}

func (x *message) Decode (s Stream) {
  x.byte = s[0]
  x.s.Decode (s[1:])
}
