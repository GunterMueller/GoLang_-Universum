package msghs

// (c) Christian Maurer   v. 171217 - license see µU.go

import
  . "µU/obj"
type
  message struct {
                 byte "Type"
           value,
             num,
          maxnum uint
                 bool "reply ok"
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
         x.value == y.value &&
         x.num == y.num &&
         x.maxnum == y.maxnum &&
         x.bool == y.bool
}

func (x *message) Copy (Y Any) {
  y := x.imp(Y)
  x.byte = y.byte
  x.value = y.value
  x.num = y.num
  x.maxnum = y.maxnum
  x.bool = y.bool
}

func (x *message) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *message) Codelen() uint {
  return 1 + 3 * C0 + 1
}

func (x *message) Encode() []byte {
  bs := make([]byte, x.Codelen())
  bs[0] = x.byte
  i, a := uint(1), C0
  copy(bs[i:i+a], Encode(x.value))
  i += a
  copy(bs[i:i+a], Encode(x.num))
  i += a
  copy(bs[i:i+a], Encode(x.maxnum))
  i += a
  bs[i] = 0; if x.bool { bs[i] = 1 }
  return bs
}

func (x *message) Decode (bs []byte) {
  x.byte = bs[0]
  i, a := uint(1), C0
  x.value = Decode(uint(0), bs[i:i+a]).(uint)
  i += a
  x.num = Decode(uint(0), bs[i:i+a]).(uint)
  i += a
  x.maxnum = Decode(uint(0), bs[i:i+a]).(uint)
  i += a
  x.bool = bs[i] == 1
}

func (x *message) Type() byte {
  return x.byte
}

func (x *message) Val() uint {
  return x.value
}

func (x *message) Num() uint {
  return x.num
}

func (x *message) Maxnum() uint {
  return x.maxnum
}

func (x *message) Ok() bool {
  return x.bool
}

func (x *message) Reply (t bool) {
  x.bool = t
}

func (x *message) Pass (t byte, v, n, m uint) {
  x.byte = t
  x.value, x.num, x.maxnum = v, n, m
}

func (x *message) Define (t byte, v uint) {
  x.byte, x.value = t, v
}
