package langs

// (c) Christian Maurer   v. 170919 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
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
