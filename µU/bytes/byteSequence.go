package bytes

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"
type
  byteSequence struct {
                      Stream
                      }

func new_(n uint) ByteSequence {
  return &byteSequence { make (Stream, n) }
}

func (x *byteSequence) imp (Y any) *byteSequence {
  y, ok := Y.(*byteSequence)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *byteSequence) Empty() bool {
  for _, a := range (x.Stream) {
    if a != byte(0) {
      return false
    }
  }
  return true
}

func (x *byteSequence) Clr() {
  for i := 0; i < len (x.Stream); i++ {
    x.Stream[i] = byte(0)
  }
}

func (x *byteSequence) Copy (Y any) {
  y := x.imp (Y)
  if len(y.Stream) != len (x.Stream) { return }
  copy (x.Stream, y.Stream)
}

func (x *byteSequence) Clone() any {
  y := new_(uint(len (x.Stream))).(*byteSequence)
  copy (y.Stream, x.Stream)
  return y
}

func (x *byteSequence) Eq (Y any) bool {
  y := x.imp (Y)
  if len (y.Stream) != len (x.Stream) { return false }
  for i, a := range (y.Stream) {
    if x.Stream[i] != a {
      return false
    }
  }
  return true
}

func (x *byteSequence) Less (Y any) bool {
  return false // TODO lexicographic ? ? ?
}

func (x *byteSequence) Leq (Y any) bool {
  return false
}

func (x *byteSequence) Codelen() uint {
  return uint(len (x.Stream))
}

func (x *byteSequence) Encode() Stream {
  b := make (Stream, len (x.Stream))
  copy (b, x.Stream)
  return b
}

func (x *byteSequence) Decode (b Stream) {
  if len (b) == len (x.Stream) {
    copy (x.Stream, b)
  } else {
    x.Clr()
  }
}
