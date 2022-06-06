package obj

// (c) Christian Maurer   v. 220420 - license see µU.go

type
  Editor interface { // Objects, that can be written to a particular
                     // position of a screen and that can be changed
                     // by interaction with a user (e.g. by pressing
                     // keys on a keyboard or a mouse).
                     //
                     // A position on a screen is given by line- or
                     // pixeloriented coordinates, i.e., by pairs of
                     // unsigned integers (l, c) or integers (x, y),
                     // where l = line and c = column on the screen,
                     // x = pixel in horizontal and y = pixel in
                     // vertical direction. In both cases (0, 0)
                     // denotes the top left corner of the screen.
  Object

// Pre: l, c have to be "small enough", i.e.
//      l + height of x < scr.NLines, c + width of x < scr.NColums.
// x is written to the screen with
// its left top corner at line/column = l/c.
  Write (l, c uint)

// Pre: see Write.
// x has the value, that was edited at line/column l/c.
// Hint: A "new" object is "read" by editing an empty one.
  Edit (l, c uint)
}

func IsEditor (a any) bool {
  if a == nil { return false }
  _, ok := a.(Editor)
  return ok
}

type
  EditorGr interface {

  Editor

// Pre: see above. x, y are pixel coordinates.
  WriteGr (x, y int)
  EditGr (x, y int)
}

func IsEditorGr (a any) bool {
  if a == nil { return false }
  _, ok := a.(EditorGr)
  return ok
}
