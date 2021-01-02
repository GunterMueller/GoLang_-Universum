package xwin

// (c) Christian Maurer   v. 201230 - license see µU.go

// #include <stdio.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
/*
unsigned long xGetPixel (XImage *i, int x, int y) { return XGetPixel (i, x, y); }
void xPutPixel (XImage *i, int x, int y, unsigned long p) { XPutPixel (i, x, y, p); }
void xDestroyImage (XImage *i) { XDestroyImage (i); }
*/
import
  "C"
import
  "µU/obj"

func enc (x uint32) obj.Stream {
  s := make(obj.Stream, 4)
  for i := 0; i < 4; i++ {
    s[i] = byte(x)
    x >>= 8
  }
  s[0], s[2] = s[2], s[0]
  return s
}

func (X *xwindow) Codelen (w, h uint) uint {
  return 4 * 4 + 4 * w * h
}

func (X *xwindow) Encode (x, y, w, h uint) obj.Stream {
  if w == 0 || h == 0 { panic ("xwin.Encode: w == 0 or h == 0") }
  if w > uint(X.wd) { panic ("xwin.Encode: w > X.wd") }
  if h > uint(X.ht) { panic ("xwin.Encode: h > X.ht") }
  s := make (obj.Stream, X.Codelen (w, h))
  n := 4 * 4
  copy (s[:n], obj.Encode4 (uint32(x), uint32(y), uint32(w), uint32(h)))
  ximg := C.XGetImage (dpy, C.Drawable(X.win), C.int(x), C.int(y), C.uint(w), C.uint(h),
                       C.ulong(1 << 24 - 1), C.XYPixmap)
  var pixel uint32
  for y := 0; y < int(h); y++ {
    for x := 0; x < int(w); x++ {
      pixel = uint32(C.xGetPixel (ximg, C.int(x), C.int(y)))
      copy (s[n:n+4], enc (pixel))
      n += 4
    }
  }
  C.xDestroyImage (ximg)
  return s
}

func (X *xwindow) Decode (s obj.Stream) {
  if s == nil { return }
  n := int32(4 * 4) // !
  x, y, w, h := obj.Decode4 (s[:n])
  ximg := C.XGetImage (dpy, C.Drawable(X.win), C.int(x), C.int(y),
                       C.uint(w), C.uint(h), C.ulong(1 << 24 - 1), C.XYPixmap)
  var pixel C.ulong
  for j := uint32(0); j < h; j++ {
    for i := uint32(0); i < w; i++ {
      pixel = (C.ulong)(obj.Decode (n, s[n:n+4]).(int32)) // (uint32)
      C.xPutPixel (ximg, C.int(i), C.int(j), pixel)
      n += 4
    }
  }
  C.XPutImage (dpy, C.Drawable(X.win), X.gc, ximg, 0, 0, C.int(x), C.int(y), C.uint(w), C.uint(h))
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.buffer), X.gc, C.int(x), C.int(y),
               C.uint(w), C.uint(h), C.int(x), C.int(y))
  C.xDestroyImage (ximg)
  C.XFlush (dpy)
}
