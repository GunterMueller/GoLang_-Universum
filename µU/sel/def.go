package sel

// (c) Christian Maurer   v. 191125 - license see µU.go

import (
  "µU/col"
  "µU/font"
)
const (
  M = 16
  N = 36
)
type // TODO description of semantics !
  WritingCol func (uint, uint, uint, col.Colour, col.Colour)

// Pre: 1 < n => m > 0; Z < scr.NLines - 2; w >= 1, c + w <= scr.NColumns; i < n.
// The calling process was blocked, until user has selected a value i <= n with keyboard or mouse.
// Until then a bar menue of height min(n, scr.NY1 - l - 1) and width w
// was written to the screen, starting at line l, column c,
// consisting of at most h texts >>> TODO <<< with fore/background colour b/f,
// one of which has inverted colours (at the beginning this is t[i]).
// Either user has chosen one of the texts with arrow keys or mouse
// (then i < n is now the number, that corresponds to the selected text)
// or she has cancelled the selection (then now i == n).
// The bar menue now has disappeared from the screen and its place on the screen is restored.
func Select (wc WritingCol, n, h, w uint, i *uint, l, c uint, f, b col.Colour) { select_(wc,n,h,w,i,l,c,f,b) }

// Pre: 1 < n <= HIGH (T) + 1; l < scr.NLines - 2; w >= 1, c + w <= scr.NColumns; i < n.
// The calling process was blocked, until user has selected a value i <= n with keyboard or mouse.
// Until then a bar menue of height min(N, scr.NY1 - l - 1) and width w
// was written to the screen, starting at line l, column c,
// consisting of h of the texts t[i] with fore-/background colour b/f,
// one of which has inverted colours (at the beginning this is t[i]).
// Either user has chosen one of the texts with arrow keys or mouse
// (then i < n is now the number, that corresponds to the selected text)
// or she has cancelled the selection (then now i == n).
// The bar menue now has disappeared from the screen and its place on the screen is restored.
func Select1 (t []string, n, w uint, i *uint, l, c uint, f, b col.Colour) { select1(t,n,w,i,l,c,f,b) }

// >>>  experimental version.
// Pre: m is part of a mask; s is a non empty file-suffix;
//      w > 0 (width); l < scr.NLines, c + i < scr.NColumns.
// Returns name + "." + string selected by user.
func Names (m, s string, i uint, l, c uint, f, b col.Colour) (string, string) { return names(m,s,i,l,c,f,b) }

// Returns an interactively selected Colour.
func Colour16() col.Colour { return colour16() }
func Colour() col.Colour { return colour() }

// Returns an interactively selected font size.
func Size (c col.Colour) font.Size { return size(c) }
