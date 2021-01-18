package xwin

// (c) Christian Maurer   v. 210105 - license see µU.go

// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
/*
unsigned long xGetPixel (XImage *i, int x, int y) { return ((*((i)->f.get_pixel))((i), (x), (y))); }
void xPutPixel (XImage *i, int x, int y, unsigned long p) { ((*((i)->f.put_pixel))((i), (x), (y), (p))); }
void xDestroyImage (XImage *i) { ((*((i)->f.destroy_image))((i))); }
*/
import
  "C"
import (
  "µU/obj"
  "µU/col"
)
const
  M = C.ulong(1 << 32 - 1)

func (X *xwindow) Codelen (w, h uint) uint {
  return 2 * 4 + 3 * w * h
}

func (X *xwindow) Encode (x0, y0, w, h uint) obj.Stream {
  if w == 0 || h == 0 { panic ("xwin.Encode: w == 0 or h == 0") }
  if w > uint(X.wd) { panic ("xwin.Encode: w > X.wd") }
  if h > uint(X.ht) { panic ("xwin.Encode: h > X.ht") }
  s := make (obj.Stream, X.Codelen (w, h))
  i := 2 * 4
  copy (s[:i], obj.Encode4 (uint16(x0), uint16(y0), uint16(w), uint16(h)))
  ximg := C.XGetImage (dpy, C.Drawable(X.win), C.int(x0), C.int(y0),
                       C.uint(w), C.uint(h), M, C.XYPixmap)
  var pixel C.ulong
  for y := 0; y < int(h); y++ {
    for x := 0; x < int(w); x++ {
      pixel = C.xGetPixel (ximg, C.int(x), C.int(y))
      e := obj.Encode(uint32(pixel))
      copy (s[i:i+3], e)
      s[i], s[i+2] = s[i+2], s[i]
      i += 3
    }
  }
  C.xDestroyImage (ximg)
  return s
}

func (X *xwindow) Decode (s obj.Stream) {
  if s == nil { return }
  n := uint32(2 * 4)
  x0, y0, w, h := obj.Decode4 (s[:n])
  ximg := C.XGetImage (dpy, C.Drawable(X.win), C.int(x0), C.int(y0),
                       C.uint(w), C.uint(h), M, C.XYPixmap)
  var pixel C.ulong
  c := col.New()
  for j := uint16(0); j < uint16(h); j++ {
    for i := uint16(0); i < uint16(w); i++ {
      c.Set (s[n+2], s[n+1], s[n+0])
      pixel = (C.ulong)(c.Code())
      C.xPutPixel (ximg, C.int(i), C.int(j), pixel)
      n += 3
    }
  }
  C.XPutImage (dpy, C.Drawable(X.win), X.gc, ximg, 0, 0, C.int(x0), C.int(y0),
               C.uint(w), C.uint(h))
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.buffer), X.gc, C.int(x0), C.int(y0),
               C.uint(w), C.uint(h), C.int(x0), C.int(y0))
  C.XFlush (dpy)
}
