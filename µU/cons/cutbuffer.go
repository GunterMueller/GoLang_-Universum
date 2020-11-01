package cons

// (c) Christian Maurer   v. 201016 - license see µU.go

import
  . "µU/obj"
var
  buffer Stream

func (x *console) Copy (s string) {
  buffer = make(Stream, len(s))
  copy (buffer[:], Stream(s))
}

func (x *console) Paste() string {
  return string(buffer[:])
}
