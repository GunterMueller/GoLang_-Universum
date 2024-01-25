package internal

// (c) Christian Maurer   v. 231220 - license see µU.go

import
  . "µU/obj"
type
  message struct {
                 byte "message kind"
                 uint "identity"
             num,
          maxnum uint
                 bool "reply ok"
}

func new_() Message {
  return &message{ Candidate, 0, 0, 0, false }
}

func (x *message) imp(Y any) *message {
  y, ok := Y.(*message)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *message) Eq (Y any) bool {
  y := x.imp(Y)
  return x.byte == y.byte && x.uint == y.uint &&
         x.num == y.num && x.maxnum == y.maxnum &&
         x.bool == y.bool
}

func (x *message) Copy (Y any) {
  y := x.imp(Y)
  x.byte, x.uint = y.byte, y.uint
  x.num, x.maxnum = y.num, y.maxnum
  x.bool = y.bool
}

func (x *message) Clone() any {
  y := new_()
  y.Copy(x)
  return y
}

func (x *message) Codelen() uint {
  return 1 + 3 * C0 + 1
}

func (x *message) Encode() Stream {
  s := make(Stream, x.Codelen())
  s[0] = x.byte
  i := uint(1)
  copy(s[i:i+C0], Encode(x.uint))
  i += C0
  copy(s[i:i+C0], Encode(x.num))
  i += C0
  copy(s[i:i+C0], Encode(x.maxnum))
  i += C0
  s[i] = 0; if x.bool { s[i] = 1 }
  return s
}

func (x *message) Decode (s Stream) {
  x.byte = s[0]
  i := uint(1)
  x.uint = Decode(uint(0), s[i:i+C0]).(uint)
  i += C0
  x.num = Decode(uint(0), s[i:i+C0]).(uint)
  i += C0
  x.maxnum = Decode(uint(0), s[i:i+C0]).(uint)
  i += C0
  x.bool = s[i] == 1
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
