package sel

// (c) Christian Maurer   v. 220815 - license see µU.go

import (
  "µU/col"
  "µU/font"
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
func Select (wc WritingCol, n, h, w uint, i *uint, l, c uint, f, b col.Colour) {
  select_(wc, n, h, w, i, l, c, f, b)
}

// Pre: 1 < n <= len (T) + 1; l < scr.NLines - 2; w >= 1, c + w <= scr.NColumns; i < n.
// The calling process was blocked, until user has selected a value i <= n with keyboard or mouse.
// Until then a bar menue of height min(N, scr.NY1 - l - 1) and width w
// was written to the screen, starting at line l, column c,
// consisting of h of the texts t[i] with fore-/background colour b/f,
// one of which has inverted colours (at the beginning this is t[i]).
// Either user has chosen one of the texts with arrow keys or mouse
// (then i < n is now the number, that corresponds to the selected text)
// or she has cancelled the selection (then now i == n).
// The bar menue now has disappeared from the screen and its place on the screen is restored.
func Select1 (t []string, n, w uint, i *uint, l, c uint, f, b col.Colour) {
  select1 (t, n, w, i, l, c, f, b)
}

// Returns an interactively selected Colour, true;
// returns Black, false), if the selection was cancelled.
//func Colour (x, y int) (col.Colour, bool) { return colour(x,y) } // 28 colours

func Colour (l, c, w uint) (col.Colour, bool) { return colour(l,c,w) }
func Colours (l, c, w uint, cols ...col.Colour) (col.Colour, bool) { return colours(l,c,w,cols...) }

// Returns an interactively selected font size.
func Fontsize (f, b col.Colour) font.Size { return fontsize(f,b) }
