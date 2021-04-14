package atom

// (c) Christian Maurer   v. 210410 - license see µU.go

import (
  "reflect"
  "µU/str"
  . "µU/obj"
  "µU/ker"
  "µU/font"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/enum"
  "µU/tval"
  "µU/text"
  "µU/bn"
  "µU/br"
  "µU/clk"
  "µU/day"
  "µU/euro"
  "µU/pers"
  "µU/phone"
  "µU/addr"
  "µU/cntry"
  "µU/sel"
)
type
  atom struct {
              Object
              uint
         w, h uint
              }
var
  name = []string {"Enumerator ",
                   "TruthValue ",
                   "Text       ",
                   "Natural    ",
                   "Real       ",
                   "Clocktime  ",
                   "Calendarday",
                   "Euro       ",
                   "Person     ",
                   "PhoneNumber",
                   "Address    ",
                   "Country    "}

func new_(o Object) Atom {
  x := new(atom)
  if o == nil {
    x.Object, x.uint = nil, 0
    return x
  }
  s := reflect.TypeOf (o).String()
  if p, ok := str.Pos (s, '.'); ok {
    s = s[1:p]
  }
  x.h, x.w = uint(1), uint(0)
  switch s {
  case "enum":
    x.Object, x.uint = enum.New (o.(enum.Enumerator).Typ()), Enumerator
    x.w = o.(enum.Enumerator).Wd()
  case "tval":
    x.Object, x.uint = tval.New(), TruthValue
    x.w = 1
  case "text":
    x.w = o.(text.Text).Len()
    x.Object, x.uint = text.New (x.w), Text
  case "bn":
    x.w = o.(bn.Natural).Width()
    x.Object, x.uint = bn.New (x.w), Natural
  case "br":
    x.w = o.(br.Real).Width()
    x.Object, x.uint = br.New (x.w), Real
  case "clk":
    x.Object, x.uint = clk.New(), Clocktime
    x.w = 5
  case "day":
    x.Object, x.uint = day.New(), Calendarday
    x.w = 10
  case "euro":
    x.Object, x.uint = euro.New(), Euro
    x.w = 10
  case "pers":
    x.Object, x.uint = pers.New(), Person
    x.h, x.w = 2, 80
  case "phone":
    x.Object, x.uint = phone.New(), PhoneNumber
    x.w = 16
  case "addr":
    x.Object, x.uint = addr.New(), Address
    x.h, x.w = 3, 80
  case "cntry":
    x.Object, x.uint = cntry.New(), Country
    x.w = 22
  default:
    ker.Panic ("atom.New: parameter object is not of an admissible type")
  }
  return x
}

func (x *atom) imp(Y Any) *atom {
  y, ok := Y.(*atom)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *atom) Type() uint {
  return x.uint
}

func (x *atom) Obj() Object {
  return x.Object
}

func (x *atom) Size() (uint, uint) {
  return x.h, x.w
}

func (x *atom) Equiv (Y Any) bool {
  return x.uint == x.imp (Y).uint
}

func (x *atom) Empty() bool {
  return x.Object.Empty()
}

func (x *atom) Clr() {
  x.Object.Clr()
}

func (x *atom) Eq (Y Any) bool {
  return x.Object.Eq (x.imp (Y).Object)
}

func (x *atom) Copy (Y Any) {
  x.Object.Copy (x.imp (Y).Object)
}

func (x *atom) Clone() Any {
  y := new_(x.Object)
  y.Copy(x)
  return y
}

func (x *atom) Less (Y Any) bool {
  return x.Object.Less (x.imp (Y).Object)
}

func (x *atom) GetFormat() Format {
  return x.Object.(Formatter).GetFormat()
}

func (x *atom) SetFormat (f Format) {
  x.Object.(Formatter).SetFormat (f)
}

func (x *atom) Colours (f, b col.Colour) {
  x.Object.(col.Colourer).Colours (f, b)
}

func (x *atom) Write (l, c uint) {
  x.Object.(Editor).Write (l, c)
}

func (x *atom) Edit (l, c uint) {
  x.Object.(Editor).Edit (l, c)
}

func (x *atom) SetFont (f font.Font) {
  x.Object.(Printer).SetFont (f)
}

func (x *atom) Print (l, c uint) {
  x.Object.(Printer).Print (l, c)
}

func (x *atom) Codelen() uint {
  return x.Object.Codelen()
}

func (x *atom) Encode() Stream {
  return x.Object.Encode()
}

func (x *atom) Decode (bs Stream) {
  x.Object.Decode (bs)
}

func (x *atom) String() string {
  return x.Object.(Stringer).String()
}

func (x *atom) Defined (s string) bool {
  return x.Object.(Stringer).Defined (s)
}

func (x *atom) Rotate() {
  x.Object.(Rotator).Rotate()
}

func (x *atom) Selected (l, c uint) bool {
  cF, cB := scr.ScrCols()
  cB.Contrast()
  n := uint(0)
  w := uint(len(name[0]))
  z, s := scr.MousePos()
  sel.Select1 (name, uint(Ntypes), w, &n, z, s, cF, cB)
  if n < Ntypes {
    x.uint = n
  } else {
    return false
  }
  x.h = 1
  b := bn.New(2)
  scr.Colours (col.White(), col.Black())
  switch x.uint {
  case Enumerator:
    e := enum.Title // TODO select e
/*/
    name = []string {"Title      ",
                     "AudioC     ".
                     "BookC      ".
                     "Writer     ".
                     "Composer   ".
                     "AudioMedium".
                     "Religion   "}
/*/
    x.Object = enum.New (e)
    n = 10 // XXX
    x.w = n
  case TruthValue:
    x.Object = tval.New()
    x.w = 1
  case Text:
    errh.Hint ("bitte die Länge des Textes eingeben")
    b.Edit (l, c)
//    scr.Write ("  ", l, c)
    errh.DelHint()
    n = b.Val()
    x.Object = text.New (n)
    x.w = n
  case Natural:
    errh.Hint ("bitte die Anzahl der Stellen der Zahl eingeben")
    b.Edit (l, c)
    scr.Write ("  ", l, c)
    errh.DelHint()
    n = b.Val()
    x.Object = bn.New (n)
    x.w = n
  case Real:
    errh.Hint ("bitte die Anzahl der Vorkommastellen der Zahl eingeben")
    b.Edit (l, c)
    scr.Write ("  ", l, c)
    errh.DelHint()
    n = b.Val()
    x.Object = br.New (n)
    x.w = n + 4
  case Clocktime:
    x.Object = clk.New()
    x.w = 5
  case Calendarday:
    x.Object = day.New()
    x.w = 10
  case Euro:
    x.Object = euro.New()
    x.w = 10
  case Person:
    x.Object = pers.New()
    x.w, x.h = 80, 2
  case PhoneNumber:
    x.Object = phone.New()
    x.w = 16
  case Address:
    x.Object = addr.New()
    x.w, x.h = 80, 3
  case Country:
    x.Object = cntry.New()
    x.w = 22
  }
  return true
}
