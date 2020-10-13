package xwin

// (c) Christian Maurer   v. 200904 - license see µU.go

// #include <stdlib.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/cursorfont.h>
import
  "C"
import (
  "unsafe"
  "strconv"
  "µU/font"
)

func (X *xwindow) ActFontsize() font.Size {
  return X.fontsize
}

func (X *xwindow) SetFontsize (s font.Size) {
  X.fontsize = s
  name := "-xos4-terminus-bold"
  h := int(font.Ht (s))
  if s < font.Normal {
    name = "-misc-fixed-medium"
  }
  name += "-r-*-*-" + strconv.Itoa(h) + "-*-*-*-*-*-*-*"
  f := C.CString (name); defer C.free (unsafe.Pointer(f))
  if dpy == nil { panic ("xwin.SetFontsize: dpy == nil") }
  X.fsp = C.XLoadQueryFont (dpy, f)
  if X.fsp == nil { panic ("terminus-fonts are not installed !") }
  X.ht1 = uint(h)
  X.wd1, X.bl1 = uint(X.fsp.max_bounds.width), uint(X.fsp.max_bounds.ascent)
  if X.bl1 + uint(X.fsp.max_bounds.descent) != X.ht1 { panic ("xwin: font bl + d != ht") }
  C.XSetFont (dpy, X.gc, C.Font(X.fsp.fid))
}

func (X *xwindow) Wd1() uint {
  return uint(X.wd1)
}

func (X *xwindow) Ht1() uint {
  return X.ht1
}

func (X *xwindow) NLines() uint {
  return uint(X.ht / X.ht1)
}

func (X *xwindow) NColumns() uint {
  return uint(X.wd / X.wd1)
}
