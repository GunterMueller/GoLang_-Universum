package atom

// (c) Christian Maurer   v. 230326 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/str"
  "µU/box"
  "µU/errh"
  "µU/sel"
  "µU/N"
  "µU/bn"
  "µU/br"
  "µU/text"
  "µU/day"
  "µU/clk"
  "µU/euro"
  "µU/phone"
  "µU/cntry"
  "µU/enum"
)
const
  M = 64 // maximal string length
type
  atom struct {
              int // typ
              text.Text
              bn.Natural
              br.Real
              day.Calendarday
              clk.Clocktime
              euro.Euro
              phone.PhoneNumber
              cntry.Country
              enum.Enum
              string // name for Enum
      l, c, w uint
         f, b col.Colour
              bool // is index
              }
var (
  w = []string {"Zeichenkette ",
                "natürl. Zahl ",
                "reelle Zahl  ",
                "Datum        ",
                "Uhrzeit      ",
                "Geldbetrag   ",
                "Telefonnummer",
                "Staat        ",
                "Aufzähltyp   "}
  wlen = uint(len(w[0]))
  bx = box.New()
)

func new_() Atom {
  x := new (atom)
  x.int = String
  x.f, x.b = col.FlashWhite(), col.Blue()
  x.w = 1
  return x
}

func (x *atom) imp (Y any) *atom {
  y, ok := Y.(*atom)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *atom) Place (l, c uint) {
  x.l, x.c = l, c
}

func (x *atom) Pos() (uint, uint) {
  return x.l, x.c
}

func (x *atom) Width() uint {
  return x.w
}

func (x *atom) Index (b bool) {
  x.bool = b
}

func (x *atom) EditIndex() {
  errh.Hint ("(Umschalt-)Enter: (kein) Index")
  for {
    switch k, d := kbd.Command(); k {
    case kbd.Enter:
      x.bool = d == 0
      errh.DelHint()
      return
    }
  }
}

func (x *atom) IsIndex() bool {
  return x.bool
}

func z (t string, max, l uint) uint {
  scr.Clr (l, 0, scr.NColumns(), 1)
  c := uint(len(t))
  bx.Colours (col.FlashWhite(), col.Black())
  bx.Colours (col.Black(), col.FlashWhite())
  bx.Wd (c)
  bx.Write (t, l, 0)
  w := uint(2); if max < 10 { w = uint(1) }
  bx.Wd (w)
  var t1 string
  for {
    t1 = str.New (w)
    bx.Edit (&t1, l, c + 1)
    if cmd, _ := kbd.LastCommand(); cmd != kbd.Go {
      nn, ok := N.Natural (t1)
      if ok {
        if nn == 0 {
          errh.Error0 ("Zahl muss > 0 sein")
        } else {
          if nn <= max {
            n0 := scr.NColumns()
            bx.Wd (n0)
            bx.Colours (col.FlashWhite(), col.Black())
            bx.Write (str.New (n0), l, 0)
            return nn
          }
          errh.Errorm ("Zahl muss <=", max, "sein")
        }
      } else {
        errh.Error0 ("keine Zahl")
      }
    }
  }
}

func (x *atom) SelectColF() {
  if c, ok := sel.Colours (x.l, x.c, x.w, col.FlashWhite(), col.Yellow()); ok {
    x.f.Copy (c)
  }
}

func (x *atom) SelectColB() {
  if c, ok := sel.Colours (x.l, x.c, x.w, col.Brown(), col.Red(), col.DarkGreen(),
                           col.DarkCyan(), col.Blue(), col.DarkGray()); ok {
    x.b.Copy (c)
  }
}

func (x *atom) EnumName (n string) {
  if x.int != Enum { ker.Panic ("x has not type Enum") }
  if str.Empty(n) { ker.Panic (n + " is empty") }
  x.string = n
}

func (x *atom) EnumSet (l, c uint) {
  if x.int != Enum { ker.Panic ("x has not type Enum") }
  if str.Empty (x.string) { ker.Panic ("name for Enum is empty") }
  x.Enum.SetEdit (x.string, l, c)
}

func (x *atom) EnumGet() {
  if x.int != Enum { ker.Panic ("x has not type Enum") }
  if str.Empty (x.string) { ker.Panic ("name for Enum is empty") }
  x.Enum.Get (x.string)
}

func (x *atom) Select() {
  n := uint(0)
  l, c := scr.MousePos()
  for {
    sel.Select1 (w, Ntypes, wlen, &n, l, c, col.Black(), col.FlashWhite())
    if n < Ntypes {
      break
    }
  }
  x.int = int(n)
  l = scr.NLines() - 1
  switch x.int {
  case String:
    x.w = z (str.Lat1("maximale Länge:"), M, l)
    x.Text = text.New (x.w)
  case Natural:
    x.w = z ("maximale Stellenzahl:", 10, l)
    x.Natural = bn.New (x.w)
  case Real:
    d := z ("Anzahl Vorkommastellen:", 10, l)
    x.Real = br.New (d)
    x.w = d + 4
  case Calendarday:
    x.w = 8
    x.Calendarday = day.New()
  case Clocktime:
    x.w = 5
    x.Clocktime = clk.New()
  case Euro:
    x.w = 10
    x.Euro = euro.New()
  case PhoneNumber:
    x.w = 16
    x.PhoneNumber = phone.New()
  case Country:
    x.w = 22
    x.Country = cntry.New()
  case Enum:
    x.w = z (str.Lat1("maximale Länge:"), enum.M, l)
    x.Enum = enum.New (x.w)
  }
}

func (x *atom) Typ() int {
  return x.int
}

func (x *atom) Define (t int, w uint) {
  if t >= Ntypes { ker.PrePanic() }
  x.int = t
  x.w = w
  switch t {
  case String:
    x.Text = text.New(x.w)
  case Natural:
    x.Natural = bn.New (x.w)
  case Real:
    x.Real = br.New (x.w - 4)
  case Calendarday:
    x.Calendarday = day.New()
  case Clocktime:
    x.Clocktime = clk.New()
  case Euro:
    x.Euro = euro.New()
  case PhoneNumber:
    x.PhoneNumber = phone.New()
  case Country:
    x.Country = cntry.New()
  case Enum:
    x.Enum = enum.New (x.w)
  }
}

func (x *atom) Empty() bool {
  switch x.int {
  case String:
    return x.Text.Empty()
  case Natural:
    return x.Natural.Empty()
  case Real:
    return x.Real.Empty()
  case Calendarday:
    return x.Calendarday.Empty()
  case Clocktime:
    return x.Clocktime.Empty()
  case Euro:
    return x.Euro.Empty()
  case PhoneNumber:
    return x.PhoneNumber.Empty()
  case Country:
    return x.Country.Empty()
  case Enum:
    return x.Enum.Empty()
  }
  return false
}

func (x *atom) Clr() {
  switch x.int {
  case String:
    x.Text.Clr()
  case Natural:
    x.Natural.Clr()
  case Real:
    x.Real.Clr()
  case Calendarday:
    x.Calendarday.Clr()
  case Clocktime:
    x.Clocktime.Clr()
  case Euro:
    x.Euro.Clr()
  case PhoneNumber:
    x.PhoneNumber.Clr()
  case Country:
    x.Country.Clr()
  case Enum:
    x.Enum.Clr()
  }
}

func (x *atom) Copy (Y any) {
  y := x.imp (Y)
  x.int = y.int
  x.l, x.c, x.w = y.l, y.c, y.w
  x.f.Copy (y.f)
  x.b.Copy (y.b)
  x.bool = y.bool
  switch y.int {
  case String:
    x.Text = text.New (y.Text.Len())
    x.Text.Copy (y.Text)
  case Natural:
    x.Natural = bn.New (y.Natural.Width())
    x.Natural.Copy (y.Natural)
  case Real:
    x.Real = br.New (y.Real.Width() - 4)
    x.Real.Copy (y.Natural)
  case Calendarday:
    x.Calendarday = day.New()
    x.Calendarday.Copy (y.Calendarday)
  case Clocktime:
    x.Clocktime = clk.New()
    x.Clocktime.Copy (y.Clocktime)
  case Euro:
    x.Euro.Copy (y.Euro)
    x.Euro.Copy (y.Euro)
  case PhoneNumber:
    x.PhoneNumber = phone.New()
    x.PhoneNumber.Copy (y.PhoneNumber)
  case Country:
    x.Country = cntry.New()
    x.Country.Copy (y.Country)
  case Enum:
    x.Enum = enum.New (x.w)
    x.Enum.Copy (y.Enum)
  }
}

func (x *atom) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *atom) Eq (Y any) bool {
  y := x.imp (Y)
  if y.int != x.int { return false }
  switch x.int {
  case String:
    return x.Text.Eq (y.Text)
  case Natural:
    return x.Natural.Eq (y.Natural)
  case Real:
    return x.Real.Eq (y.Real)
  case Calendarday:
    return x.Calendarday.Eq (y.Calendarday)
  case Clocktime:
    return x.Clocktime.Eq (y.Clocktime)
  case Euro:
    return x.Euro.Eq (y.Euro)
  case PhoneNumber:
    return x.PhoneNumber.Eq (y.PhoneNumber)
  case Country:
    return x.Country.Eq (y.Country)
  case Enum:
    return x.Enum.Eq (y.Enum)
  }
  return false
}

func (x *atom) Less (Y any) bool {
  y := x.imp (Y)
  if y.int != x.int { return false }
  switch x.int {
  case String:
    return x.Text.Less (y.Text)
  case Natural:
    return x.Natural.Less (y.Natural)
  case Real:
    return x.Real.Less (y.Real)
  case Clocktime:
    return x.Clocktime.Less (y.Clocktime)
  case Euro:
    return x.Euro.Less (y.Euro)
  case PhoneNumber:
    return x.PhoneNumber.Less (y.PhoneNumber)
  case Country:
    return x.Country.Less (y.Country)
  case Enum:
    return x.Enum.Less (y.Enum)
  }
  return false
}

func (x *atom) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *atom) Sub (Y any) bool {
  y := x.imp (Y)
  if y.int != x.int { return false }
  switch x.int {
  case String:
    _, ok := str.Sub (x.Text.String(), y.Text.String())
    return ok
  case Natural:
    return x.Natural.Eq (y.Natural)
  case Real:
    return x.Real.Eq (y.Real)
  case Clocktime:
    return x.Clocktime.Eq (y.Clocktime)
  case Euro:
    return x.Euro.Eq (y.Euro)
  case PhoneNumber:
    return x.PhoneNumber.Eq (y.PhoneNumber)
  case Country:
    return x.Country.Eq (y.Country)
  case Enum:
    return x.Enum.Eq (y.Enum)
  }
  return false
}

func (x *atom) PosLess (Y any) bool {
  y := x.imp (Y)
  l, c := x.Pos()
  l1, c1 := y.Pos()
  if l == l1 { return c < c1 }
  return l < l1
}

func (x *atom) Colours (f, b col.Colour) {
  x.f.Copy (f)
  x.b.Copy (b)
  switch x.int {
  case String:
    x.Text.Colours (f, b)
  case Natural:
    x.Natural.Colours (f, b)
  case Real:
    x.Real.Colours (f, b)
  case Calendarday:
    x.Calendarday.Colours (f, b)
  case Clocktime:
    x.Clocktime.Colours (f, b)
  case Euro:
    x.Euro.Colours (f, b)
  case PhoneNumber:
    x.PhoneNumber.Colours (f, b)
  case Country:
    x.Country.Colours (f, b)
  case Enum:
    x.Enum.Colours (f, b)
  }
}

func (x *atom) Cols() (col.Colour, col.Colour) {
  return x.f, x.b
}

func (x *atom) Write (l, c uint) {
  x.Colours (x.f, x.b)
  l, c = x.l, x.c
  switch x.int {
  case String:
    x.Text.Write (l, c)
  case Natural:
    x.Natural.Write (l, c)
  case Real:
    x.Real.Write (l, c)
  case Calendarday:
    x.Calendarday.Write (l, c)
  case Clocktime:
    x.Clocktime.Write (l, c)
  case Euro:
    x.Euro.Write (l, c)
  case PhoneNumber:
    x.PhoneNumber.Write (l, c)
  case Country:
    x.Country.Write (l, c)
  case Enum:
    x.Enum.Write (l, c)
  }
}

func (x *atom) Edit (l, c uint) {
  l, c = x.l, x.c
  x.Write (l, c)
  switch x.int {
  case String:
    x.Text.Edit (l, c)
  case Natural:
    x.Natural.Edit (l, c)
  case Real:
    x.Real.Edit (l, c)
  case Calendarday:
    x.Calendarday.Edit (l, c)
  case Clocktime:
    x.Clocktime.Edit (l, c)
  case Euro:
    x.Euro.Edit (l, c)
  case PhoneNumber:
    x.PhoneNumber.Edit (l, c)
  case Country:
    x.Country.Edit (l, c)
  case Enum:
    x.Enum.Edit (l, c)
  }
}

func (x *atom) Print (l, c uint) {
  switch x.int {
  case String:
    x.Text.Print (l, c)
  case Natural:
    x.Natural.Print (l, c)
  case Real:
    x.Real.Print (l, c)
  case Calendarday:
    x.Calendarday.Print (l, c)
  case Clocktime:
    x.Clocktime.Print (l, c)
  case Euro:
    x.Euro.Print (l, c)
  case PhoneNumber:
    x.PhoneNumber.Print (l, c)
  case Country:
    x.Country.Print (l, c)
  case Enum:
    x.Enum.Print (l, c)
  }
}

func (x *atom) Codelen() uint {
  switch x.int {
  case String:
    return x.Text.Codelen()
  case Natural:
    return x.Natural.Codelen()
  case Real:
    return x.Real.Codelen()
  case Calendarday:
    return x.Calendarday.Codelen()
  case Clocktime:
    return x.Clocktime.Codelen()
  case Euro:
    return x.Euro.Codelen()
  case PhoneNumber:
    return x.PhoneNumber.Codelen()
  case Country:
    return x.Country.Codelen()
  case Enum:
    return x.Enum.Codelen()
  }
  return 0
}

func (x *atom) Encode() Stream {
  s := make(Stream, x.Codelen())
  switch x.int {
  case String:
    a := x.Text.Codelen()
    copy (s[0:a], x.Text.Encode())
  case Natural:
    a := x.Natural.Codelen()
    copy (s[0:a], x.Natural.Encode())
  case Real:
    a := x.Real.Codelen()
    copy (s[0:a], x.Real.Encode())
  case Calendarday:
    a := x.Calendarday.Codelen()
    copy (s[0:a], x.Calendarday.Encode())
  case Clocktime:
    a := x.Clocktime.Codelen()
    copy (s[0:a], x.Clocktime.Encode())
  case Euro:
    a := x.Euro.Codelen()
    copy (s[0:a], x.Euro.Encode())
  case PhoneNumber:
    a := x.PhoneNumber.Codelen()
    copy (s[0:a], x.PhoneNumber.Encode())
  case Country:
    a := x.Country.Codelen()
    copy (s[0:a], x.Country.Encode())
  case Enum:
    a := x.Enum.Codelen()
    copy (s[0:a], x.Enum.Encode())
  }
  return s
}

func (x *atom) Decode (s Stream) {
  switch x.int {
  case String:
    x.Text = text.New (x.w)
    a := x.Text.Codelen()
    x.Text.Decode (s[0:a])
  case Natural:
    x.Natural = bn.New (x.w)
    a := x.Natural.Codelen()
    x.Natural.Decode (s[0:a])
  case Real:
    x.Real = br.New (x.w - 4)
    a := x.Real.Codelen()
    x.Real.Decode (s[0:a])
  case Calendarday:
    x.Calendarday = day.New()
    a := x.Calendarday.Codelen()
    x.Calendarday.Decode (s[0:a])
  case Clocktime:
    x.Clocktime = clk.New()
    a := x.Clocktime.Codelen()
    x.Clocktime.Decode (s[0:a])
  case Euro:
    x.Euro = euro.New()
    a := x.Euro.Codelen()
    x.Euro.Decode (s[0:a])
  case PhoneNumber:
    x.PhoneNumber = phone.New()
    a := x.PhoneNumber.Codelen()
    x.PhoneNumber.Decode (s[0:a])
  case Country:
    x.Country = cntry.New()
    a := x.Country.Codelen()
    x.Country.Decode (s[0:a])
  case Enum:
    x.Enum = enum.New (x.w)
    a := x.Enum.Codelen()
    x.Enum.Decode (s[0:a])
  }
}

func (x *atom) String() string {
  switch x.int {
  case String:
    return x.Text.String()
  case Natural:
    return x.Natural.String()
  case Real:
    return x.Real.String()
  case Calendarday:
    return x.Calendarday.String()
  case Clocktime:
    return x.Clocktime.String()
  case Euro:
    return x.Euro.String()
  case PhoneNumber:
    return x.PhoneNumber.String()
  case Country:
    return x.Country.String()
  case Enum:
    return x.Enum.String()
  }
  return ""
}
