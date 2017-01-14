package bytes

// (c) murus.org  v. 140330 - license see murus.go

// >>> Just for fun, most likely completely worthless

import
  . "murus/obj"
type
  byteSequence struct {
                    s []byte
                      }

func (x *byteSequence) imp (Y Any) *byteSequence {
  y, ok:= Y.(*byteSequence)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func New (n uint) ByteSequence {
  return &byteSequence { make ([]byte, n) }
}

func (x *byteSequence) Empty() bool {
  for _, a:= range (x.s) {
    if a != byte(0) {
      return false
    }
  }
  return true
}

func (x *byteSequence) Clr() {
  for i:= 0; i < len (x.s); i++ {
    x.s[i] = byte(0)
  }
}

func (x *byteSequence) Copy (Y Any) {
  y:= x.imp (Y)
  if len (y.s) != len (x.s) { return }
  Copy (x.s, y.s)
}

func (x *byteSequence) Clone() Any {
  y:= New (uint(len (x.s)))
  Copy (y.s, x.s)
  return y
}

func (x *byteSequence) Eq (Y Any) bool {
  y:= x.imp (Y)
  if len (y.s) != len (x.s) { return false }
  for i, a:= range (y.s) {
    if x.s[i] != a {
      return false
    }
  }
  return true
}

func (x *byteSequence) Less (Y Any) bool {
  return false // TODO lexicographic ? ? ?
}

func (x *byteSequence) Codelen() uint {
  return uint(len (x.s))
}

func (x *byteSequence) Encode() []byte {
  b:= make ([]byte, len (x.s))
  copy (b, x.s)
  return b
}

func (x *byteSequence) Decode (b []byte) {
  if len (b) == len (x.s) {
    copy (x.s, b)
  } else {
    x.Clr()
  }
}
