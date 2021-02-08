package internal

// (c) Christian Maurer   v. 201014 - license see µU.go

import . "nU/obj"

type message struct {
  byte "message type"
  uint "identity"
  num, maxnum uint
  bool "reply ok"
}
var c0 uint

func init() {
  c0 = C0()
}

func new_() Message {
  return &message{ Candidate, 0, 0, 0, false }
}

func (x *message) imp(Y Any) *message {
  y, ok := Y.(*message)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *message) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.byte == y.byte &&
         x.uint == y.uint &&
         x.num == y.num &&
         x.maxnum == y.maxnum &&
         x.bool == y.bool
}

func (x *message) Copy (Y Any) {
  y := x.imp(Y)
  x.byte = y.byte
  x.uint = y.uint
  x.num, x.maxnum = y.num, y.maxnum
  x.bool = y.bool
}

func (x *message) Clone() Any {
  y := new_()
  y.Copy(x)
  return y
}

func (x *message) Codelen() uint {
  return 1 + 3 * c0 + 1
}

func (x *message) Encode() Stream {
  bs := make(Stream, x.Codelen())
  bs[0] = x.byte
  i, a := uint(1), c0
  copy(bs[i:i+a], Encode(x.uint))
  i += a
  copy(bs[i:i+a], Encode(x.num))
  i += a
  copy(bs[i:i+a], Encode(x.maxnum))
  i += a
  bs[i] = 0; if x.bool { bs[i] = 1 }
  return bs
}

func (x *message) Decode (bs Stream) {
  x.byte = bs[0]
  i, a := uint(1), c0
  x.uint = Decode(uint(0), bs[i:i+a]).(uint)
  i += a
  x.num = Decode(uint(0), bs[i:i+a]).(uint)
  i += a
  x.maxnum = Decode(uint(0), bs[i:i+a]).(uint)
  i += a
  x.bool = bs[i] == 1
}

func (x *message) Kind() byte {
  return x.byte
}

func (x *message) IdNumsOk() (uint, uint, uint, bool) {
  return x.uint, x.num, x.maxnum, x.bool
}

func (x *message) SetPass (v, n, m uint) {
  x.byte = Candidate
  x.uint, x.num, x.maxnum = v, n, m
}

func (x *message) SetReply (b bool) {
  x.byte = Reply
  x.bool = b
}

func (x *message) SetLeader (v uint) {
  x.byte = Leader
  x.uint = v
}
