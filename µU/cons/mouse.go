package cons

// (c) Christian Maurer   v. 170918 - license see µU.go

import (
  "µU/mouse"
  "µU/ptr"
  "µU/col"
)

func (X *console) MouseEx() bool {
  return mouse.Ex()
}

func (X *console) MousePos() (uint, uint) {
  xm, ym := mouse.Pos()
  return uint(ym - X.y) / X.ht1, uint(xm - X.x) / X.wd1 // Offset
}

func (X *console) MousePosGr() (int, int) {
  xm, ym := mouse.Pos()
  return xm - X.x, ym - X.y // Offset
}

func (X *console) SetPointer (p ptr.Pointer) {
  X.pointer = p
}

var (
  pointer = [ptr.NPointers][]string {
    []string { // Ptr
      "# . . . . . . . . . . . . ",
      "# # . . . . . . . . . . . ",
      "# o # . . . . . . . . . . ",
      "# * o # . . . . . . . . . ",
      "# * * o # . . . . . . . . ",
      "# * * * o # . . . . . . . ",
      "# * * * * o # . . . . . . ",
      "# * * * * * o # . . . . . ",
      "# * * * * * * o # . . . . ",
      "# * * * * * * * o # . . . ",
      "# * * * * * * * * o # . . ",
      "# * * * * * * * * * o # . ",
      "# * * * * * # # # # # # # ",
      "# * * * o * o # . . . . . ",
      "# * * # * * * # . . . . . ",
      "# * # # # o * o # . . . . ",
      "# # . . . # * * # . . . . ",
      "# . . . . # o * o . . . . ",
      ". . . . . . # o # . . . . " },
    []string { // Gumby
      ". . # # # # # # . . . . . . . . ",
      "* * o # * * * * # . . . . . . . ",
      "# # * o # * * * * # . . . . . . ",
      "# # # * # * . * . * # . . . . . ",
      "# # * o # * * * * * # . . . . . ",
      "# # * o # * . . . * # * * * . . ",
      "# # # # # * * * * * # # # # * . ",
      "* * # # # * * * * * # # # # # # ",
      ". . * * # * * * * * # . * # # # ",
      ". . . . # * * * * * # . * # # # ",
      ". . . . # * * # * * # * # # # # ",
      ". . . . # * * # * * # . * # # # ",
      ". . . . # * * # * * # . . * * * ",
      ". . . # * * * # * * * # . . . . ",
      ". . # * * * * # * * * * # . . . ",
      ". . # # # # # . # # # # # . . . " },
    []string { // Hand
      ". . . o # # o . . . . . . . . . ",
      ". . . # * * # . . . . . . . . . ",
      ". . . # * * # . . . . . . . . . ",
      ". . . # * * # . . . . . . . . . ",
      ". . . # * * # # # o . . . . . . ",
      ". . . # * * # o * o # # o . . . ",
      ". . . # * * # * * # o * # # # o ",
      ". # # # * * # * * * * * # o * # ",
      "# o * # * * # * * * * * * * * # ",
      "# * * # * * * * * * * * * * * # ",
      "# * * # * * * * * * * * * * * # ",
      "# * * * * o o * o o * o o * * # ",
      "# * * * * * * * * * * * * * * # ",
      "# * * * * * * * * * * * * * * # ",
      "# o * * * * * * * * * * * * o # ",
      ". # # # # # # # # # # # # # # . " },
    []string { // Gobbler
      ". . . . . . . . * * * * * * . . ",
      ". . . . . . . . * # # # # * . . ",
      "* * . . . . . . * # # # * * * * ",
      "# * * * * * * * * * # # * * # # ",
      "# * * # # # # # # * # # * * * * ",
      "# # # # # # # # # # # # * * . . ",
      "# # # # # # # * * * # # * * . . ",
      "# # # # # # * * * * # # * * . . ",
      "* # # * * * * * * * # # # * . . ",
      "* * * * * * * * # # # # * * . . ",
      ". * * # * # # # # # # * * . . . ",
      ". . . * # * * * * * * * . . . . ",
      ". . . * # * . . . . . . . . . . ",
      ". . . * # * . . . . . . . . . . ",
      ". . * * # * * * . . . . . . . . ",
      ". . * # # # # * . . . . . . . . " },
    []string { // Watch
      ". . . . o # # # # # o . . . . . ",
      ". . . # o * * * * * # # . . . . ",
      ". . # * * * * * * * * * # . . . ",
      ". # * * * * * * * * * * * # . . ",
      "o o * * * * * * * * * * * o o . ",
      "# * * * * * * * * * * * * * # . ",
      "# * * * * * o # o * * * * * # . ",
      "# * * o # # # # # * * * * * # . ",
      "# * # # o * o # o * * * * * # . ",
      "# * * * * * * * # * * * * * # . ",
      "o o * * * * * * # * * * * o o . ",
      ". # * * * * * * * # * * * # . . ",
      ". . # * * * * * * # * * # . . . ",
      ". . . # o * * * * * # # . . . . ",
      ". . . . o # # # # # o . . . . . ",
      ". . . . . . . . . . . . . . . . " },
  }
  pointerHt = [ptr.NPointers]int { 18, 16, 16, 16, 16 }
  pointerWd = [ptr.NPointers]int { 12, 16, 16, 16, 16 }
)

func (X *console) initMouse () {
//  X.mouseOn = false
  mouse.Def (uint(0), uint(0), width, height)
  X.xMouse, X.yMouse = int(X.wd) / 2, int(X.ht) / 2
  mouse.Warp (uint(X.xMouse), uint(X.yMouse))
}

func (X *console) restore (x, y int) {
  a := (int(width) * y + x) * int(colourdepth)
  da := pointerWd[X.pointer] * int(colourdepth)
  w := width * colourdepth
// TODO limit to right screen border ???
  h1, ht := pointerHt[X.pointer], int(X.ht)
  if y + h1 > ht { h1 = ht - y }
  for h := 0; h < h1; h++ {
    copy (fbmem[a:a+da], fbcop[a:a+da])
    a += int(w)
  }
}

func (X *console) writePointer (x, y int) {
  cB, cW, cG := col.Black().Cc(), col.LightWhite().Cc(), col.LightGray().Cc()
  var p []byte
  for h := 0; h < pointerHt[X.pointer]; h++ {
    for w := 0; w < pointerWd[X.pointer]; w++ {
      switch pointer[X.pointer][h][2 * w] { case '#':
        p = cB
      case '*':
        p = cW
      case 'o':
        p = cG
      default:
        continue
      }
      if x + w < X.x || x + w >= X.x + int(X.wd) ||
         y + h < X.y || y + h >= X.y + int(X.ht) {
        continue
      }
      a := (int(width) * (y + h) + (x + w)) * int(colourdepth)
      copy (fbmem[a:a+int(colourdepth)], p)
//      copy (fbcop[a:a+int(colourdepth)], p) // No, we don't want that
    }
  }
}

func (X *console) MousePointer (on bool) {
  X.mouseOn = on
  if ! mouse.Ex() || ! X.mouseOn || ! visible { return }
  if X.x <= X.xMouse + pointerWd[X.pointer] && X.xMouse < X.x + int(X.wd) &&
     X.y <= X.yMouse + pointerHt[X.pointer] && X.yMouse < X.y + int(X.ht) {
    X.restore (X.xMouse, X.yMouse)
  }
  if X == mouseConsole {
//    X.restore (X.xMouse, X.yMouse)
    X.xMouse, X.yMouse = mouse.Pos()
    X.writePointer (X.xMouse, X.yMouse)
  } else {
    // TODO full screen as root window
  }
}

func (X *console) MousePointerOn() bool {
  if ! mouse.Ex() { return false }
  return X.mouseOn
}

func (X *console) WarpMouse (l, c uint) {
  mouse.Warp (uint(X.y) + l * X.ht, uint(X.x) + c * X.wd) // Offset
  X.MousePointer (true)
}

func (X *console) WarpMouseGr (x, y int) {
  mouse.Warp (uint(x + X.x), uint(y + X.y)) // Offset
  X.MousePointer (true)
}

func (X *console) UnderMouse (l, c, w, h uint) bool {
  if ! mouse.Ex() { return false }
  lm, cm := X.MousePos()
  return l <= lm && lm < l + h && c <= cm && cm < c + w
}

func (X *console) UnderMouseGr (x, y, x1, y1 int, t uint) bool {
  if ! mouse.Ex() { return false }
  intord (&x, &y, &x1, &y1)
  xm, ym := X.MousePosGr()
  return x <= int(xm) + int(t) && int(xm) <= x1 + int(t) &&
         y <= int(ym) + int(t) && int(ym) <= y1 + int(t)
}

func (X *console) UnderMouse1 (x, y int, d uint) bool {
  if ! mouse.Ex() { return false }
  xm, ym := X.MousePosGr()
  return (x - xm) * (x - xm) + (y - ym) * (y - ym) <= int(d * d)
}
