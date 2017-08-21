package schol

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/obj"
//  "murus/enum"
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

  String() string
  FullAged() bool
  Equiv (y Any) bool
  Edit0 (l, c uint)
  Printer
  Indexer
}

func New() Scholar { return new_() }
