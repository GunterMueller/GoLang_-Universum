package obj

// (c) Christian Maurer   v. 141129 - license see murus.go

import
  "murus/col"
type
  Editor interface { // Objects, that can be printed to a particular
                     // position of a screen and that can be changed
                     // by interaction with a user (e.g. by pressing
                     // keys on a keyboard or a mouse).
                     //
                     // A position on a screen is given by logical or
                     // pixeloriented coordinates, i.e. by pairs of
                     // unsigned integers (l, c) or integers (x, y),
                     // where l = line and c = column on the screen,
                     // x = pixel in horizontal and y = pixel in
                     // vertical direction. In both cases (0, 0)
                     // denotes the top left corner of the screen.
                     //
                     // So particularly any Editor is an
  Object             // (details see murus/obj/object.go).

// x has the colours fore-/background f/b.
  Colours (f, b col.Colour)

// Pre: l, c have to be "small enough", i.e.
//      l + height (object) < scr.NoLines,
//      c + width (object) < scr.NoColums.
// x is in its colours written to the screen
// with its left top corner at line, column = l, c.
  Write (l, c uint)

// Pre: see Write.
// x has the value, that was edited at line/column l/c.
// Hint: A "new" object is "read" by editing an empty one.
  Edit (l, c uint)

// >>>  eventually new version:

// x has the colours given by the parameters.
//  Colours (... col.Colour)

// Pre: If there are position parameters p[0], p[1],
//      then they have to be "small enough", i.e.
//      p[0] + height (object) < scr.NoLines,
//      p[1] + width (object) < scr.NoColums.
// x is in its colours written to the screen
// [ with its left top corner at line/column p[0]/p[1]) ].
//  Write (... uint)

// Precondition: see Write.
// x has the value, that was edited
// [ at line/column p[0]/p[1] (see Write) ].
// Hint: A "new" object is "read" by editing an empty one.
//  Edit (... uint)
}

type
  EditorGr interface {

  Editor

  WriteGr (x, y int)
  EditGr (x, y int)
}
