package langs

// (c) murus.org  v. 130115 - license see murus.go

import (
  . "murus/obj"
//  "murus/enum"
)
type
  LanguageSequence interface {

  Formatter
  Editor
  Printer

//  Num (l[]subject.Subject, v, b[]uint) uint
}
