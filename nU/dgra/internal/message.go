package internal

// (c) Christian Maurer   v. 170429 - license see nU.go

import . "nU/obj"

type message struct {
  typ MsgType
  value, num, maxnum uint
  replyOk bool
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
  return x.typ == y.typ &&
         x.value == y.value &&
         x.num == y.num &&
         x.maxnum == y.maxnum &&
         x.replyOk == y.replyOk
}

func (x *message) Copy (Y Any) {
  y := x.imp(Y)
  x.typ, x.value, x.num, x.maxnum, x.replyOk = y.typ, y.value, y.num, y.maxnum, y.replyOk
}

func (x *message) Clone() Any {
  y := new_()
  y.Copy(x)
  return y
}

func (x *message) Codelen() uint {
  return 1 + 3 * C0 + 1
}

func (x *message) Encode() []byte {
  bs := make([]byte, x.Codelen())
  bs[0] = byte(x.typ)
  i, a := uint(1), C0
  copy(bs[i:i+a], Encode(x.value))
  i += a
  copy(bs[i:i+a], Encode(x.num))
  i += a
  copy(bs[i:i+a], Encode(x.maxnum))
  i += a
  bs[i] = 0; if x.replyOk { bs[i] = 1 }
  return bs
}

func (x *message) Decode (bs []byte) {
  x.typ = MsgType(bs[0])
  i, a := uint(1), C0
  x.value = Decode(uint(0), bs[i:i+a]).(uint)
  i += a
  x.num = Decode(uint(0), bs[i:i+a]).(uint)
  i += a
  x.maxnum = Decode(uint(0), bs[i:i+a]).(uint)
  i += a
  x.replyOk = bs[i] == 1
}

func (x *message) Typ() MsgType {
  return x.typ
}

func (x *message) Content() (uint, uint, uint, bool) {
  return x.value, x.num, x.maxnum, x.replyOk
}

func (x *message) Val() uint {
  return x.value
}

func (x *message) Reply (t bool) {
  x.typ = Reply
  x.replyOk = t
}

func (x *message) PassCandidate (i, n, m uint) {
  x.typ = Candidate
  x.value, x.num, x.maxnum = i, n, m
}

func (x *message) Define (t MsgType, v uint){
  x.typ, x.value = t, v
}
