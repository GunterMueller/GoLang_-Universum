package xker

// (c) murus.org  v. 140615 - license see murus.go

// #include <stdlib.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
import
  "C"

func (X *window) clr (x, y int, w, h uint) {
  C.XSetForeground (dpy, X.gc, cc (X.scrB))
  C.XFillRectangle (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h))
  C.XFillRectangle (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h))
  C.XSetForeground (dpy, X.gc, cc (X.scrF))
  C.XFlush (dpy)
}

func (X *window) Clr (l, c, w, h uint) {
  x, y:= int(c) * int(X.wd1), int(l) * int(X.ht1)
  x1, y1:= int(c + w) * int(X.wd1), int(l + h) * int(X.ht1)
  X.ClrGr (x, y, x1, y1)
}

func (X *window) ClrGr (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  X.clr (x, y, uint(x1 - x) + 1, uint(y1 - y) + 1) // incl. right and bottom border; man XDrawRectangle
}

func (X *window) Cls() {
  X.clr (0, 0, uint(X.wd), uint(X.ht))
}

func (X *window) Buffered () bool {
  return X.buff
}

func (X *window) Buf (on bool) {
  if X.buff == on { return }
  X.buff = on
  if on {
    C.XSetForeground (dpy, X.gc, cc (X.scrB))
    C.XFillRectangle (dpy, C.Drawable(X.buffer), X.gc, nul, nul, X.wd, X.ht)
    C.XSetForeground (dpy, X.gc, cc (X.scrF))
    C.XFlush (dpy)
  } else {
    X.buf2win()
  }
}

func natord (x, y, x1, y1 *uint) {
  if *x > *x1 { *x, *x1 = *x1, *x }
  if *y > *y1 { *y, *y1 = *y1, *y }
}

func (X *window) Save (l, c, w, h uint) {
  x, y:= C.int(X.wd1) * C.int(c), C.int(X.ht1) * C.int(l)
  w_, h_:= X.wd1 * C.uint(w), X.ht1 * C.uint(h)
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.shadow), X.gc, x, y, w_, h_, x, y)
}

func (X *window) SaveGr (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  w, h:= C.uint(x1 - x + 1), C.uint(y1 - y + 1)
//  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.shadow), X.gc, C.int(x), C.int(y), w, h, C.int(x), C.int(y))
  C.XCopyArea (dpy, C.Drawable(X.buffer), C.Drawable(X.shadow), X.gc, C.int(x), C.int(y), w, h, C.int(x), C.int(y))
}

func (X *window) Save1() {
  X.SaveGr (0, 0, int(X.wd) - 1, int(X.ht) - 1)
}

func (X *window) Restore (l, c, w, h uint) {
  x, y:= C.int(X.wd1) * C.int(c), C.int(X.ht1) * C.int(l)
  w_, h_:= X.wd1 * C.uint(w), X.ht1 * C.uint(h)
  C.XCopyArea (dpy, C.Drawable(X.shadow), C.Drawable(X.win), X.gc, x, y, w_, h_, x, y)
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.buffer), X.gc, x, y, w_, h_, x, y)
}

func (X *window) RestoreGr (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  w, h:= C.uint(x1 - x + 1), C.uint(y1 - y + 1)
  C.XCopyArea (dpy, C.Drawable(X.shadow), C.Drawable(X.win), X.gc, C.int(x), C.int(y), w, h, C.int(x), C.int(y))
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.buffer), X.gc, C.int(x), C.int(y), w, h, C.int(x), C.int(y))
}

func (X *window) Restore1() {
  X.RestoreGr (0, 0, int(X.wd) - 1, int(X.ht) - 1)
}
