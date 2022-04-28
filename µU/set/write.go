package set

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "µU/scr"
  "µU/str"
)

func (x *node) write (x0, x1, y, dy uint, f func (any) string) {
  if x == nil { return }
  xm := (x0 + x1) / 2
  y1 := int(y + scr.Ht1() / 2) - 1
  if x.left != nil {
    scr.Line (int(xm), y1, int(x0 + xm) / 2, y1 + int (dy))
  }
  if x.right != nil {
    scr.Line (int(xm), y1, int(xm + x1) / 2, y1 + int (dy))
  }
  scr.WriteGr (f (x.any), int(xm - scr.Wd1()), int(y))
  x.left.write (x0, xm, y + dy, dy, f)
  x.right.write (xm, x1, y + dy, dy, f)
}

func (x *set) Write (x0, x1, y, dy uint, f func (any) string) {
  x.anchor.write (x0, x1, y, dy, f)
}

func (x *node) write1 (d uint, f func (any) string) {
  if x == nil { return }
  x.right.write1 (d + 1, f)
  println (str.New(6 * d) + f (x.any))
  x.left.write1 (d + 1, f)
}

func (x *set) Write1 (f func (any) string) {
  x.anchor.write1 (0, f)
}
