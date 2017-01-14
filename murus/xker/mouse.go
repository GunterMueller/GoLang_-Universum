package xker

// (c) murus.org  v. 160109 - license see murus.go

// #include <X11/X.h>
// #include <X11/Xlib.h>
import
  "C"
import
  "murus/ptr"

func (X *window) MouseEx() bool {
  return true // a mouse should be running on any GUI
}

func (X *window) SetPointer (p ptr.Pointer) {
  C.XDefineCursor (dpy, X.win, C.XCreateFontCursor (dpy, C.uint(ptr.Code (p))))
}

func (X *window) MousePos() (uint, uint) {
  return uint(X.yM) / uint(X.ht1), uint(X.xM) / uint(X.wd1)
}

func (X *window) MousePosGr() (int, int) {
  return X.xM, X.yM
}

func (X *window) WarpMouse (l, c uint) {
  X.WarpMouseGr (int(c) * int(X.wd1), int(l) * int(X.ht1))
}

func (X *window) WarpMouseGr (x, y int) {
  C.XWarpPointer (dpy, C.None, X.win, nul, nul, unul, unul, C.int(x), C.int(y))
  C.XFlush (dpy)
}

func (X *window) MousePointer (b bool) {
  X.mousePointer (b)
}

func (X *window) MousePointerOn() bool {
  return X.mousePointerOn()
}

func (X *window) UnderMouse (l, c, w, h uint) bool {
  xm, ym:= X.MousePos()
  return l <= xm && xm < l + h && c <= ym && ym < c + w
}

func (X *window) UnderMouseGr (x, y, x1, y1 int, t uint) bool {
  intord (&x, &y, &x1, &y1)
  xm, ym:= X.MousePosGr()
  return x <= xm + int(t) && xm <= x1 + int(t) && y <= ym + int(t) && ym <= y1 + int(t)
}

func (X *window) UnderMouse1 (x, y int, d uint) bool {
  xm, ym:= X.MousePosGr()
  return (x - xm) * (x - xm) + (y - ym) * (y - ym) <= int(d * d)
}
