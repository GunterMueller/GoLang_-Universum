package xker

// (c) murus.org  v. 140615 - license see murus.go

// #include <X11/X.h>
// #include <X11/Xlib.h>
/*
// both functions do not return the top left coordinates of the window, but simply 0 !
int x (Display *d, Drawable w) { XWindowAttributes a; XGetWindowAttributes (d, w, &a); return a.x; }
int y (Display *d, Drawable w) { XWindowAttributes a; XGetWindowAttributes (d, w, &a); return a.y; }
*/
import
  "C"
import
  . "murus/mode"

func maxMode() Mode {
  initX()
  return fullScreen
}

func maxRes() (uint, uint) {
  initX()
  return uint(width), uint(height)
}

func ok (m Mode) bool {
  fullScreen = maxMode()
  return Wd(m) <= Wd(fullScreen) && Ht(m) <= Ht(fullScreen)
}

func (X *window) X() uint {
  return uint(C.x (dpy, C.Drawable(X.win))) // does not work - see above remark !
  return uint(X.x) // always the initial value
}

func (X *window) Y() uint {
  return uint(C.y (dpy, C.Drawable(X.win))) // does not work - see above remark !
  return uint(X.y) // always the initial value
}

func (X *window) Wd() uint {
  return uint(X.wd)
}

func (X *window) Ht() uint {
  return uint(X.ht)
}

func (X *window) Wd1() uint {
  return uint(X.wd1)
}

func (X *window) Ht1() uint {
  return uint(X.ht1)
}

func (X *window) NLines() uint {
  return uint(X.ht / X.ht1)
}

func (X *window) NColumns() uint {
  return uint(X.wd / X.wd1)
}
