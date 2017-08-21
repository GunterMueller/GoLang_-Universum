package xker

// (c) murus.org  v. 140615 - license see murus.go

// #include <stdlib.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/cursorfont.h>
import
  "C"
import (
  "unsafe"; "strconv"
  "murus/font"
)

func (X *window) ActFontsize() font.Size {
  return X.fontsize
}

// /usr/share/fonts/misc
func (X *window) SetFontsize (s font.Size) {
  X.fontsize = s
  const (mf = "misc-fixed"; te = "terminus")
  name:= "-"
  switch s { case font.Tiny, font.Small:
    name += mf + "-medium"
  case font.Normal, font.Big, font.Huge:
    name += "xos4-" + te + "-bold"
  default:
    return
  }
  h:= int(font.Ht (s))
  name += "-r-*-*-" + strconv.Itoa(h) + "-*-*-*-*-*-iso8859-15"
  f:= C.CString (name); defer C.free (unsafe.Pointer(f))
  if dpy == nil { panic ("xker.SetFontheight: dpy == nil") }
  X.fsp = C.XLoadQueryFont (dpy, f)
  if X.fsp == nil {
    if s <= font.Small { name = mf } else { name = te }
    panic (name + "-font is not installed !")
  }
  X.ht1 = C.uint(h)
  X.wd1, X.bl1 = C.uint(X.fsp.max_bounds.width), C.uint(X.fsp.max_bounds.ascent)
  if X.bl1 + C.uint(X.fsp.max_bounds.descent) != X.ht1 { panic ("xker: font bl + d != ht") }
//  C.XSetFont (dpy, X.gc, C.Font(X.fsp.fid))
//  X.nLines, X.nColumns = X.ht / X.ht1, X.wd / X.wd1
}
