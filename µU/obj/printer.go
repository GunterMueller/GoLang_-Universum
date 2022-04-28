package obj

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/font"
type
  Printer interface {

// f is the actual font.
  SetFont (f Font)

// Pre: x consists only of strings.
// x is printed starting in linc l, column c.
  Print (l, c uint)
}

func IsPrinter (a any) bool {
  if a == nil { return false }
  _, ok := a.(Printer)
  return ok
}
