package pbox

// (c) Christian Maurer   v. 230112 - license see µU.go

import (
  "µU/fontsize"
  "µU/font"
)
type
  Printbox interface {

// TODO Spec
  SetFont (f font.Font)

// TODO Spec
  SetFontsize (s fontsize.Size)

// TODO Spec
  Font() font.Font

// TODO Spec
  Fontsize() fontsize.Size

// TODO Spec
  Print (s string, l, c uint)

// TODO Spec
  PageReady()
}

// Returns a new print box with font Roman.
func New() Printbox { return new_() }
