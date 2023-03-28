package attr

// (c) Christian Maurer   v. 210329 - license see µU.go

import
  . "µU/obj"
const
  Wd = 3
type
  Attr = uint
type
  Attribute interface {

  Object
  Editor
  Printer
  Stringer
  Marker
}
type
  AttrSet interface {

  Object

// a ist in die Menge eingefügt.
  Ins (a Attribute)

// Die Terminattribute aus der Menge sind im Kurzformat
// in einer Zeichenkette aufgereiht ab Position (l, c)
// auf dem Bildschirm ausgegeben. Wenn b == true ist, ist 
// an Position (l, c-1) ein rotes Leerzeichen ausgegeben.
  Write (l, c uint, b bool)
}
func NewSet() AttrSet { return newSet() }
