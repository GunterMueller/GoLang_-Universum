package char

// (c) Christian Maurer   v. 170810 - license see murus.go

import (
  . "murus/obj"
  . "murus/shape"
  "murus/z"
  "murus/kbd"
  "murus/col"
  "murus/scr"
  "murus/font"
  "murus/pbox"
)
type
  character struct {
                   byte
            cF, cB col.Colour
                   font.Font
                   }
var
  pbx = pbox.New()

func new_() Character {
  x := new (character)
  x.byte = ' '
  x.cF, x.cB = scr.StartCols()
  return x
}

func (x *character) imp (Y Any) byte {
  y, ok := Y.(*character)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y.byte
}

func (x *character) SetByte (b byte) { // TODO UTF8
  if b < ' ' { b = ' ' }
  x.byte = b
}

func (x *character) Empty() bool {
  return x.byte == ' '
}

func (x *character) Clr() {
  x.byte = ' '
}

func (x *character) Eq (Y Any) bool {
  return x.byte == x.imp (Y)
}

func (x *character) Equiv (Y Any) bool {
  return x.byte % 32 == x.imp (Y) % 32
}

func (x *character) Copy (Y Any) {
  x.byte = x.imp (Y)
}

func (x *character) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *character) Val() uint {
  return uint(x.byte)
}

func (x *character) SetVal (n uint) bool {
  if n >= 1<<8 {
    return false
  }
  b := byte (n)
  if b < ' ' { b = ' ' } // TODO ausschalten unnützer Werte
  x.byte = b
  return true
}

func (x *character) ByteVal() byte {
  return x.byte
}

func (x *character) Less (Y Any) bool {
  return x.byte < x.imp (Y)
}

func (x *character) String() string {
  return string(x.byte)
}

func (x *character) Defined (s string) bool {
  if len (s) == 0 { return false }
  if len (s) > 1 { return false } // TODO UTF8
  x.byte = byte(s[0])
  return true
}

func (x *character) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *character) Write (l, c uint) {
  scr.Colours (x.cF, x.cB)
  scr.Write1 (x.byte, l, c)
}

func (x *character) Edit (l, c uint) {
  b := x.byte
  scr.Colours (x.cF, x.cB)
  scr.Write1 (b, l, c)
  loop: for {
    scr.Warp (l, c, Understroke)
    b = kbd.Byte()
    switch { case ' ' <= b, b < 128, z.IsLatin1(b):
      break loop
    }
  }
  scr.Write1 (b, l, c)
  x.byte = b
}

func (x *character) SetFont (f font.Font) {
  x.Font = f
}

func (x *character) Print (l, c uint) {
  pbx.SetFont (x.Font)
  pbx.Print (string(x.byte), l, c)
}

func (x *character) Codelen() uint {
  return 1
}

func (x *character) Encode() []byte {
  b := make ([]byte, 1)
  b[0] = x.byte
  return b
}

func (x *character) Decode (b []byte) {
  x.byte = b[0]
}
