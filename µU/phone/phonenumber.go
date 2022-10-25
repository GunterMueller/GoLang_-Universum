package phone

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/col"
  "µU/box"
  "µU/errh"
  "µU/n"
  "µU/font"
  "µU/pbox"
)
type
  phonenumber struct {
                     uint16 "prefix"
                     uint32 "number"
              cF, cB col.Colour
                     font.Font
                     }
var (
  bx = box.New()
  pbx = pbox.New()
)

func new_() PhoneNumber {
  x := new(phonenumber)
  x.cF, x.cB = col.StartCols()
  return x
}

func (x *phonenumber) imp (Y any) *phonenumber {
  y, ok := Y.(*phonenumber)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *phonenumber) Empty() bool {
  return x.uint16 == 0 && x.uint32 == 0
}

func (x *phonenumber) Clr() {
  x.uint16, x.uint32 = 0, 0
}

func (x *phonenumber) Copy (Y any) {
  y := x.imp (Y)
  x.uint16, x.uint32 = y.uint16, y.uint32
}

func (x *phonenumber) Clone() any {
  y := new_()
  y.Copy(x)
  return y
}

func (x *phonenumber) Eq (Y any) bool {
  y := x.imp (Y)
  return x.uint16 == y.uint16 && x.uint32 == y.uint32
}

func (x *phonenumber) Less (Y any) bool {
  y := x.imp (Y)
  if x.uint16 < y.uint16 { return true }
  if x.uint16 == y.uint16 { return x.uint32 < y.uint32 }
  return false
}

func (x *phonenumber) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *phonenumber) Defined (s string) bool {
  if str.Empty (s) {
    x.Clr()
    return true
  }
  str.Move (&s, true)
  l := str.ProperLen (s)
  if i, ok := str.Pos (s, ' '); ok {
    k, ok := n.Natural (s[1:i])
    if ok && s[0] == '0' {
      x.uint16 = uint16(k)
      if l == i {
        return false
      }
      s = s[i:l]
    } else {
      return false
    }
  } else {
    x.uint16 = 0
  }
  str.OffBytes (&s, ' ')
  if s == "" {
    x.uint32 = uint32(x.uint16)
    x.uint16 = 0
    return true
  }
  if tmp, ok := n.Natural (s); ok {
    x.uint32 = uint32(tmp)
    return true
  } else {
    x.uint16 = 0
    x.uint32 = 0
  }
  return false
}

func (x *phonenumber) String() string {
  s := ""
  if x.uint16 > 0 {
    s = n.String (uint(x.uint16))
    s = "0" + s
  }
  if x.uint32 > 0 {
    t := n.String (uint(x.uint32))
    n := len (t)
    switch n {
    case 4, 5:
      t = t[0:n-2] + " " + t[n-2:]
    case 6, 7:
      t = t[0:n-4] + " " + t[n-4:n-2] + " " + t[n-2:]
    case 8, 9:
      t = t[0:n-5] + " " + t[n-5:n-3] + " " + t[n-3:]
    }
    if x.uint16 == 0 {
      s = t
    } else {
      s = s + " " + t
    }
  }
  str.Norm (&s, Width)
  return s
}

func (x *phonenumber) TeX() string {
  t := x.String()
  str.OffSpc1 (&t)
  return t
}

func (x *phonenumber) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *phonenumber) Cols() (col.Colour, col.Colour) {
  return x.cF, x.cB
}

func (x *phonenumber) Write (l, c uint) {
  bx.Wd (Width)
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String(), l, c)
}

func (x *phonenumber) Edit (l, c uint) {
  bx.Wd (Width)
  s := x.String()
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
      break
    } else {
      errh.Error0("keine Telefonnummer")
    }
  }
  x.Write (l, c)
}

func (x *phonenumber) SetFont (f font.Font) {
  x.Font = f
}

func (x *phonenumber) Print (l, c uint) {
  pbx.SetFont (x.Font)
  pbx.Print (x.String(), l, c)
}

func (x *phonenumber) Codelen() uint {
  return 2 + // Codelen (x.uint16)
         4   // Codelen (x.uint32)
}

func (x *phonenumber) Encode() Stream {
  bs := make (Stream, x.Codelen())
  copy (bs[0:2], Encode (x.uint16))
  copy (bs[2:6], Encode (x.uint32))
  return bs
}

func (x *phonenumber) Decode (bs Stream) {
  x.uint16 = Decode (x.uint16, bs[0:2]).(uint16)
  x.uint32 = Decode (x.uint32, bs[2:6]).(uint32)
}
