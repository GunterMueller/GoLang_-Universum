package word

// (c) Christian Maurer   v. 220809 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
//  "µU/font"
  "µU/text"
  "µU/pseq"
)
type
  word struct {
              text.Text
              }
var (
  cf, cb = col.Black(), col.White()
  cfa, cba = col.LightWhite(), col.Red()
  actual = new_().(*word)
  file = pseq.New (actual.Text.String())
)

func init() {
  file.Name ("Suchwort.dat")
  if file.Empty() {
    actual.Clr()
    file.Put (actual.Text.String())
  } else {
    actual.Text.Defined (file.Get().(string))
  }
}

func (x *word) imp (Y any) *word {
  y, ok := Y.(*word)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func new_() Word {
  x := new (word)
  x.Text = text.New (Wd)
  x.Colours (cf, cb)
  return x
}

func (x *word) Copy (Y any) {
  y := x.imp (Y)
  x.Text.Copy (y.Text)
}

func (x *word) Clone() any {
  y := New()
  y.Copy (x)
  return y
}

func (x *word) Eq (Y any) bool {
  return x.Text.Eq (x.imp (Y).Text)
}

func (x *word) Less (Y any) bool {
  return x.Text.Less (x.imp (Y).Text)
}

func (x *word) Ok() bool {
  _, ok := actual.Text.EquivSub (x.Text)
  return ok && ! actual.Text.Empty()
}

func (x *word) Defined (s string) bool {
  x.Text.Defined (s)
  return true
}

func (x *word) Write (l, c uint) {
  if x.Empty() {
    x.Colours (cf, cb)
  } else if x.Ok() {
    x.Colours (cfa, cba)
  } else {
    x.Colours (cf, cb)
  }
  x.Text.Write (l, c)
}

func (x *word) Edit (l, c uint) {
  x.Write (l, c)
  x.Text.Edit (l, c)
  x.Write (l, c)
}

func EditActual (l, c uint) {
  actual.Text.Edit (l, c)
  file.Seek (0)
  file.Put (actual.Text.String())
}

/*/
func (x *word) SetFont (f font.Font) {
  x.Text.SetFont (f)
}

func (x *word) Print (l, c uint) {
  x.Text.Print (l, c)
}
/*/
