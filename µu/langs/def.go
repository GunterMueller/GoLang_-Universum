package langs

// (c) Christian Maurer   v. 170919 - license see µu.go

import (
  . "µu/obj"
  "µu/col"
)
type
  LanguageSequence interface {

  Formatter
  Object
  col.Colourer
  Editor
  Printer
//  Num (l[]subject.Subject, v, b[]uint) uint
}

func New() LanguageSequence { return new_() }
