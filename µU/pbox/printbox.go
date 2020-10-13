package pbox

// (c) Christian Maurer   v. 200902 - license see µU.go

import (
  "µU/str"
  "µU/font"
  "µU/prt"
)
type
  printbox struct {
                f font.Font
                s font.Size
                  }

func NLines () uint {
  return prt.NLines()
}

func NColumns () uint {
  return prt.NColumns()
}

func new_() Printbox {
  return &printbox { font.Roman, font.Normal }
}

func (x *printbox) SetFont (f font.Font) {
  x.f = f
}

func (x *printbox) Font() font.Font {
  return x.f
}

func (x *printbox) SetFontsize (s font.Size) {
  x.s = s
}

func (x *printbox) Fontsize() font.Size {
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
