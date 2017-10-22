package schol

// (c) Christian Maurer   v. 170919 - license see µU.go

import (
  . "µU/obj"
//  "µU/enum"
)
const ( // Format
  Minimal = iota //  1 Zeile,  52 Spalten
  VeryShort      //  1 Zeile,  80 Spalten
  Short          //  2 Zeilen, 80 Spalten
  Long           // 21 Zeilen, 80 Spalten
  NFormats
)
type
  Scholar interface {

  Formatter
  Indexer
  Printer

  String() string
  FullAged() bool
  Equiv (y Any) bool
  Edit0 (l, c uint)
}

func New() Scholar { return new_() }
