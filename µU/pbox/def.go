package pbox

// (c) Christian Maurer   v. 161216 - license see µU.go

import
  "µU/font"
type
  Printbox interface {

// TODO Spec
  SetFont (f font.Font)

// TODO Spec
  Font() font.Font

// TODO Spec

  Print (s string, l, c uint)

// TODO Spec
  PageReady()
}

// Returns a new print box with font Roman.
func New() Printbox { return new_() }
