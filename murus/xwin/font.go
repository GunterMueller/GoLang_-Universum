package xwin

// (c) Christian Maurer   v. 170818 - license see murus.go

// #include <stdlib.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/cursorfont.h>
import
  "C"
import (
  "unsafe"
  "strconv"
  "murus/font"
)

func (X *xwindow) ActFontsize() font.Size {
  return X.fontsize
}

// /usr/share/fonts/misc
func (X *xwindow) SetFontsize (s font.Size) {
  X.fontsize = s
  const (mf = "misc-fixed"; te = "terminus")
  name := "-"
  switch s { case font.Tiny, font.Small:
    name += mf + "-medium"
  case font.Normal, font.Big, font.Huge:
    name += "xos4-" + te + "-bold"
  default:
    return
  }
  h := int(font.Ht (s))
//  name += "-r-*-*-" + strconv.Itoa(h) + "-*-*-*-*-*-iso8859-15" // not found in openSUSE 42.3
  name += "-r-*-*-" + strconv.Itoa(h) + "-*-*-*-*-*-*-*"
  f := C.CString (name); defer C.free (unsafe.Pointer(f))
  if dpy == nil { panic ("xwin.SetFontheight: dpy == nil") }
  X.fsp = C.XLoadQueryFont (dpy, f)
  if X.fsp == nil {
    if s <= font.Small { name = mf } else { name = te }
    panic (name + "-font is not installed !")
  }
  X.ht1 = C.uint(h)
  X.wd1, X.bl1 = C.uint(X.fsp.max_bounds.width), C.uint(X.fsp.max_bounds.ascent)
  if X.bl1 + C.uint(X.fsp.max_bounds.descent) != X.ht1 { panic ("xwin: font bl + d != ht") }
//  C.XSetFont (dpy, X.gc, C.Font(X.fsp.fid))
//  X.nLines, X.nColumns = X.ht / X.ht1, X.wd / X.wd1
}

func (X *xwindow) Wd1() uint {
  return uint(X.wd1)
}

func (X *xwindow) Ht1() uint {
  return uint(X.ht1)
}

func (X *xwindow) NLines() uint {
  return uint(X.ht / X.ht1)
}

func (X *xwindow) NColumns() uint {
  return uint(X.wd / X.wd1)
}
