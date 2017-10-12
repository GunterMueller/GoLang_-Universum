package xwin

// (c) Christian Maurer   v. 170814 - license see µu.go

// #include <X11/X.h>
// #include <X11/Xlib.h>
import
  "C"
import
  "µu/ptr"

func (X *xwindow) MouseEx() bool {
  return true // a mouse should be running on any GUI
}

func (X *xwindow) SetPointer (p ptr.Pointer) {
  C.XDefineCursor (dpy, X.win, C.XCreateFontCursor (dpy, C.uint(ptr.Code (p))))
}

func (X *xwindow) MousePos() (uint, uint) {
  return uint(X.yM) / uint(X.ht1), uint(X.xM) / uint(X.wd1)
}

func (X *xwindow) MousePosGr() (int, int) {
  return X.xM, X.yM
}

func (X *xwindow) WarpMouse (l, c uint) {
  X.WarpMouseGr (int(c) * int(X.wd1), int(l) * int(X.ht1))
}

func (X *xwindow) WarpMouseGr (x, y int) {
  C.XWarpPointer (dpy, C.None, X.win, 0, 0, 0, 0, C.int(x), C.int(y))
  C.XFlush (dpy)
}

func (X *xwindow) MousePointer (b bool) {
  X.mousePointer (b)
}

func (X *xwindow) MousePointerOn() bool {
  return X.mousePointerOn()
}

func (X *xwindow) UnderMouse (l, c, w, h uint) bool {
  xm, ym:= X.MousePos()
  return l <= xm && xm < l + h && c <= ym && ym < c + w
}

func (X *xwindow) UnderMouseGr (x, y, x1, y1 int, t uint) bool {
  intord (&x, &y, &x1, &y1)
  xm, ym:= X.MousePosGr()
  return x <= xm + int(t) && xm <= x1 + int(t) && y <= ym + int(t) && ym <= y1 + int(t)
}

func (X *xwindow) UnderMouse1 (x, y int, d uint) bool {
  xm, ym:= X.MousePosGr()
  return (x - xm) * (x - xm) + (y - ym) * (y - ym) <= int(d * d)
}
