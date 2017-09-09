package obj

// (c) Christian Maurer   v. 140102 - license see murus.go

import
  . "murus/font"
type
  Printer interface {

// f is the actual font.
  SetFont (f Font)

// 
  Print (l, c uint)
}
