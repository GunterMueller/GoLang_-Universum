package pbox

// (c) murus.org  v. 161216 - license see murus.go

import (
  "murus/str"
  "murus/font"
  "murus/prt"
)
type
  printbox struct {
                f font.Font
                  }

func NLines () uint {
  return prt.NLines()
}

func NColumns () uint {
  return prt.NColumns()
}

func new_() Printbox {
  return &printbox { font.Roman }
}

func (x *printbox) SetFont (f font.Font) {
  x.f = f
}

func (x *printbox) Font() font.Font {
  return x.f
}

func (x *printbox) Print (s string, l, c uint) {
  if l >= prt.NLines() || c >= prt.NColumns() { return }
  str.OffSpc (&s)
  if len (s) == 0 { return }
  prt.Print (s, l, c, x.f)
}

func (x *printbox) PageReady() {
  prt.GoPrint()
}
