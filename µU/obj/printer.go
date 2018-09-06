package obj

// (c) Christian Maurer   v. 180902 - license see µU.go

import
  . "µU/font"
type
  Printer interface {

// f is the actual font.
  SetFont (f Font)

// 
  Print (l, c uint)
}

func IsPrinter (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Printer)
  return ok
}
