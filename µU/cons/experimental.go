package cons

// (c) Christian Maurer   v. 140217 - license see µU.go

func (X *console) Moved (x, y int) bool {
  ok := 0 <= X.x + x                          && 0 <= X.y + y                           &&
             X.x + x + int(X.wd) < int(width) &&      X.y + y + int(X.ht) < int(height)
  if ! ok {
// println ("NOT moved by x, y ==" + i(x) + ", " + i(y))
    return false
  }
//  X.Save1()
  X.x += x
  X.y += y
// println ("moved by x, y ==" + i(x) + ", " + i(y))
  return true
// X.verschoben ausgeben
}
