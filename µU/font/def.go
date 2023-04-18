package font

// (c) Christian Maurer   v. 230112 - license see µU.go

import
  "µU/fontsize"
type
  Font byte; const ( // only for prt
  Roman = Font(iota)
  Bold
  Italic
  NFonts
)
const
  M = 6 // len names
var
  Name []string

// Returns a string of len 2, that uniquely defines f and s.
func Code (f Font, s fontsize.Size) string { return code(f,s) }

// Returns the width resp. the height of a font in size s;
// for prt in pt and for scr in px.
func Wd (s fontsize.Size) uint { return wd(s) }
func Ht (s fontsize.Size) uint { return ht(s) }
