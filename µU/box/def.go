package box

// (c) Christian Maurer   v. 220524 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Box interface { // Boxes within one line of the screen to write and edit strings.

  Stringer

// Pre: n > 0.
// x has the width n.
  Wd (n uint)

// The editor mode is changed to that of a pocket calculator.
  SetNumerical()

// See scr.Transparence.
  Transparence (t bool)

// Pre: 0 < c < width of the calling box.
// At the beginning, the cursor is in position c.
  Start (c uint)

// x has the fore- and backgroundcolour of the screen.
  ScrColours ()

// x has the fore-/backgroundcolour f/b.
  Colours (f, b col.Colour)
  ColourF (f col.Colour)
  ColourB (b col.Colour)

// x is filled with an empty string.
  Clr (l, c uint)

// Pre: l < scr.NLines,
//      c + width of x <= scr.NColumns,
//      c + len(s) <= scr.NColumns.
//      width of X == 0 or len (s) <= width of x.
// If width of X was 0, now width of x == len(s).
// s is written to the screen,
// starting at position (line, column) == (l, c) in the colours of x.
  Write (s string, l, c uint)

// Pre: y <= scr.Ht - scr.Ht1.
//      x + scr.Wd1 * width of x < scr.NColumns,
//      x + scr.Wd1 * length of s < scr.NColumns.
// Like Write, starting at pixelpos (column, line) == (x, y).
  WriteGr (s string, x, y int)

// Pre: l < scr.NLines, c + width of the calling box < scr.NColumns,
//      c + len (s) < scr.NColumns
//      width of x == 0 or length of x <= width of x.
// If width of x was 0, now width of x == len(s).
// s is now the string, that was edited starting at position (l, c).
// To correct while typing, the usual keys can be used:
// - Backspace and Del to remove characters,
//   in combination with Shift or Strg to delete all,
// - arrow keys Left/Right and Pos1/End to move inside x,
// - Ins to toggle between insert mode (underline cursor)
//   and overwrite mode (block cursor).
// The cursor starts at the beginning of x.
// If s was empty, the mode starts with insert, otherwise with overwrite.
// The calling process was blocked, until the input was terminated with
// another command (see kbd) or one of the above commands with depth > 0.
  Edit (s *string, l, c uint)

// Pre: y <= scr.Ht() - scr.Ht1(),
//      x + scr.Wd1() * width (of the calling box) < scr.NColumns(),
//      x + scr.Wd1() * len(s) < scr.NColumns().
// Like Edit, starting at pixelpos (column, line) == (x, y).
  EditGr (s *string, x, y int)
}

// Returns an new box of width 0,
// the colours of the screen and the default editor mode.
func New() Box { return new_() }
