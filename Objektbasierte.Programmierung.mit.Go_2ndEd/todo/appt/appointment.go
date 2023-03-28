package appt

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  . "µU/obj"
  "µU/font"
  "µU/kbd"
  "µU/col"
  "µU/text"
  "µU/clk"
  . "todo/help"
  "todo/word"
  "todo/attr"
)
const
  textWd = uint(80 - 5 - 1 - attr.Wd - 1 - word.Wd - 2) // 56
type
  appointment struct {
                     clk.Clocktime
                     attr.Attribute
                     word.Word
                     text.Text
               field []any
                  cl []uint
                     }
var (
  fmt Format
  cS, cH, ctF, ctB = col.Black(), col.Yellow(), col.Black(), col.LightWhite()
)

func (x *appointment) imp (Y any) *appointment {
  y, ok := Y.(*appointment)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func New() Appointment {
  x := new (appointment)
  x.Clocktime = clk.New()
  x.Attribute = attr.New()
  x.Word = word.New()
  x.Text = text.New (textWd)
  x.Text.Colours (ctF, ctB)
  x.field = []any { x.Clocktime, x.Attribute, x.Word, x.Text }
  x.cl = []uint { x.Clocktime.Codelen(), x.Attribute.Codelen(), x.Word.Codelen(), textWd }
  return x
}

func (x *appointment) Empty() bool {
  return x.Clocktime.Empty() &&
         x.Attribute.Empty() &&
         x.Word.Empty() &&
         x.Text.Empty()
}

func (x *appointment) Clr() {
  x.Clocktime.Clr()
  x.Attribute.Clr()
  x.Word.Clr()
  x.Text.Clr()
}

func (x *appointment) Copy (Y any) {
  y := x.imp (Y)
  x.Clocktime.Copy (y.Clocktime)
  x.Attribute.Copy (y.Attribute)
  x.Word.Copy (y.Word)
  x.Text.Copy (y.Text)
}

func (x *appointment) Clone() any {
  y := New()
  y.Copy (x)
  return y
}

func (x *appointment) Eq (Y any) bool {
// buggy ?
  y := x.imp (Y)
  if x.Empty() {
    return y.Empty()
  } else if y.Empty() {
    return false
  }
  return x.Clocktime.Eq (y.Clocktime) &&
         x.Attribute.Eq (y.Attribute) &&
         x.Word.Eq (y.Word) && x.Text.Eq (x.Text)
}

func (x *appointment) Less (Y any) bool {
// TODO buggy ?
  y := x.imp (Y)
  if x.Empty() {
    return false
  }
  if y.Empty() {
    return true
  }
  if x.Clocktime.Eq (y.Clocktime) {
    if x.Attribute.Eq (y.Attribute) {
      if x.Word.Eq (y.Word) {
        return x.Text.Less (y.Text)
      }
      return x.Word.Less (y.Word)
    }
    return x.Attribute.Less (y.Attribute)
  }
  return x.Clocktime.Less (y.Clocktime)
}

func (x *appointment) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *appointment) HasWord() bool {
  return x.Word.Ok()
}

func (x *appointment) Attrib() attr.Attribute {
  return x.Attribute.Clone().(attr.Attribute)
}

var (
  ac, wc, tc uint = 5 + 1, ac + attr.Wd + 1, wc + word.Wd + 1
)

func (x *appointment) SetFormat (f Format) {
  fmt = f
}

func (x *appointment) Write (l, c uint) {
  x.Clocktime.Colours (cS, cH)
  x.Clocktime.Write (l, c)
  x.Attribute.Write (l, c + ac)
  if fmt == Long {
    x.Word.Write (l, c + wc)
    x.Text.Write (l, c + tc)
  }
}

var (
  CC kbd.Comm
  DD uint
)

func (x *appointment) Edit (l, c uint) {
  x.Write (l, c)
  i := 0
  for {
    switch i {
    case 0:
      x.Clocktime.Edit (l, c)
    case 1:
      x.Attribute.Edit (l, c + ac)
    case 2:
      if fmt == Long {
        x.Word.Edit (l, c + wc)
      }
    case 3:
      if fmt == Long {
        x.Text.Edit (l, c + tc)
      }
    }
    C, D := kbd.LastCommand()
    CC, DD = C, D
    switch C {
    case kbd.Esc: // , kbd.There:
      return
    case kbd.Enter: // kbd.Here:
      if D > 0 {
        return
      }
      if i < 3 {
        i++
      } else {
        return
      }
    case kbd.Pos1, kbd.End:
      return
    case kbd.Down, kbd.Right:
      if D > 0 {
        return
      }
      if i < 3 {
        i++
      } else {
        return
      }
    case kbd.Up, kbd.Left:
      if D > 0 {
        return
      }
      if i > 0 {
        i--
      } else {
        return
      }
/*/
    case kbd.PgDown:
      ;
    case kbd.PgUp:
      ;
/*/
    case kbd.Del:
      if D > 0 {
        return
      }
      x.Clr()
      x.Write (l, c)
    case kbd.Help:
      Help()
/*/
    case kbd.Search:
      return
    case kbd.Cut, kbd.Copy, kbd.Paste:
      return
    case kbd.Print:
      return
/*/
    default:
      return
    }
  }
}

func (x *appointment) SetFont (f font.Font) {
  // TODO
}

func (x *appointment) Print (l, c uint) {
  x.Clocktime.Print (l, c)
  x.Attribute.Print (l, c + 6)
  x.Word.Print (l, c + 11)
  x.Text.Print (l, c + 21)
}

func (x *appointment) Codelen() uint {
  var c uint
  for _, n := range x.cl { c += n }
  return c
}

func (x *appointment) Encode() Stream {
  return Encodes (x.field, x.cl)
}

func (x *appointment) Decode (bs Stream) {
  Decodes (bs, x.field, x.cl)
  x.Clocktime = x.field[0].(clk.Clocktime)
  x.Attribute = x.field[1].(attr.Attribute)
  x.Word = x.field[2].(word.Word)
  x.Text = x.field[3].(text.Text)
}

func (x *appointment) String() string {
  return x.Clocktime.String() +
         x.Attribute.String() +
         x.Word.String() +
         x.Text.String()
}

func (x *appointment) Defined (s string) bool {
  return x.Clocktime.Defined (s[:5]) &&
         x.Attribute.Defined (s[5:8]) &&
         x.Word.Defined (s[8:20]) &&
         x.Text.Defined (s[20:])
}
