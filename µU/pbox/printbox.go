package pbox

// (c) Christian Maurer   v. 230112 - license see µU.go

import (
  "µU/str"
  "µU/fontsize"
  "µU/font"
  "µU/prt"
)
type
  printbox struct {
                s fontsize.Size
                f font.Font
                  }

func NLines () uint {
  return prt.NLines()
}

func NColumns () uint {
  return prt.NColumns()
}

func new_() Printbox {
  return &printbox { fontsize.Normal, font.Roman }
}

func (x *printbox) SetFont (f font.Font) {
  x.f = f
}

func (x *printbox) Font() font.Font {
  return x.f
}

func (x *printbox) SetFontsize (s fontsize.Size) {
  x.s = s
}

func (x *printbox) Fontsize() fontsize.Size {
  return x.s
}

func (x *printbox) Print (s string, l, c uint) {
  if l >= prt.NLines() || c >= prt.NColumns() { return }
  str.OffSpc (&s)
  if len (s) == 0 { return }
  prt.SetFont (x.f)
  prt.SetFontsize (x.s)
  prt.Print (s, l, c)
}

func (x *printbox) PageReady() {
  prt.GoPrint()
}
