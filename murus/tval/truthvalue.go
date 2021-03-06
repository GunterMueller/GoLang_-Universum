package tval

// (c) Christian Maurer  v. 170810 - license see murus.go

import (
  . "murus/obj"
  "murus/str"
  "murus/col"
  "murus/scr"
  "murus/box"
  "murus/errh"
  "murus/font"
  "murus/pbox"
)
const
  max = 15
type
  value byte; const (
  undecidable = iota
  falseVal
  trueVal
  nValues
)
type (
  representation [nValues]string
  truthValue struct {
                    value
                    uint "uint"
                    representation
             cF, cB col.Colour
                    font.Font
                    }
)
var (
  bx = box.New()
  pbx = pbox.New()
)

func new_() TruthValue {
  x := new (truthValue)
  x.value = undecidable
  x.uint = 1
  x.representation = [nValues]string { " ", "n", "j" }
  x.cF, x.cB = scr.StartCols()
  return x
}

func (x *truthValue) imp (Y Any) *truthValue {
  y, ok := Y.(*truthValue)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *truthValue) SetFormat (f, t string) {
  x.representation [falseVal] = f
  x.uint = str.ProperLen (x.representation[falseVal])
  if x.uint == 0 {
    x.representation = [nValues]string { " ", "n", "j" }
    return
  }
  x.representation [trueVal] = t
  n := str.ProperLen (x.representation[falseVal])
  if n > x.uint { x.uint = n }
  x.representation [undecidable] = str.Clr (x.uint)
}

func (x *truthValue) Empty() bool {
  return x.value == undecidable
}

func (x *truthValue) Clr() {
  x.value = undecidable
}

func (x *truthValue) Copy (Y Any) {
  y := x.imp(Y)
  x.value = y.value
  x.cF, x.cB = y.cF, y.cB
}

func (x *truthValue) Clone() Any {
  y := new_()
  y.Copy(x)
  return y
}

func (x *truthValue) Eq (Y Any) bool {
  return x.value == x.imp(Y).value
}

func (x *truthValue) Less (Y Any) bool {
  return x.value < x.imp(Y).value
}

func (x *truthValue) Codelen() uint {
  return uint(1)
}

func (x *truthValue) Encode() []byte {
  b := make ([]byte, 1)
  b[0] = byte(x.value)
  return b
}

func (x *truthValue) Decode (b []byte) {
  x.value = value(b[0])
}

func (x *truthValue) Defined (s string) bool {
  switch s[0] { case ' ', '?':
    x.value = undecidable
    return true
  }
  for v := value(1); v < nValues; v++ {
    p, ok := str.EquivSub (s, x.representation[v])
    if ok && p == 0 {
      x.value = v
      return true
    }
  }
  return false
}

func (x *truthValue) String() string {
  return x.representation [x.value]
}

func (x *truthValue) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *truthValue) Write (l, c uint) {
  bx.Wd (x.uint)
  bx.Colours (x.cF, x.cB)
  bx.Write (x.representation[x.value], l, c)
}

func (x *truthValue) Edit (l, c uint) {
  bx.Wd (x.uint)
  bx.Colours (x.cF, x.cB)
  var input string
  for {
    input = x.representation[x.value]
    bx.Write (input, l, c)
    bx.Edit (&input, l, c)
    if x.Defined (input) {
      break
    } else {
      errh.Error0("Eingabe unverständlich") // , l + 1, c)
    }
  }
  x.Write (l, c)
}

func (x *truthValue) SetFont (f font.Font) {
  pbx.SetFont (f)
}

func (x *truthValue) Print (l, c uint) {
  pbx.SetFont (x.Font)
  pbx.Print (x.representation[x.value], l, c)
}

func (x *truthValue) Set (b bool) {
  x.value = falseVal
  if b {
    x.value = trueVal
  }
}
