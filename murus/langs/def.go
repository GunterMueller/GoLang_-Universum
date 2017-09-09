package langs

// (c) Christian Maurer   v. 170121 - license see murus.go

import
  . "murus/obj"
type
  LanguageSequence interface {

  Formatter
  Editor
  Printer

//  Num (l[]subject.Subject, v, b[]uint) uint
}

func New() LanguageSequence { return new_() }
