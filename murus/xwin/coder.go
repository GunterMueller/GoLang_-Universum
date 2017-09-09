package xwin

// (c) Christian Maurer   v. 170814 - license see murus.go

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
import
  "murus/obj"

/*
func cd (x int32) []byte {
  return Encode (x)
  const m = int32(256)
  b := make ([]byte, 4)
  for i := 0; i < 4; i++ {
    b[i] = byte(x % m)
    x /= m
  }
  return b
}
*/

func (X *xwindow) Codelen (w, h uint) uint {
  return 4 * (4 + w * h)
}

func (X *xwindow) Encode (x, y, w, h uint) []byte {
  if w == 0 || h == 0 || w > uint(X.wd) || h > uint(X.ht) { panic ("xwin.Encode: w or h has wrong value") }
  bs := make ([]byte, X.Codelen (w, h))
  n := 4 * 4
  copy (bs[:n], obj.Encode4 (uint32(x), uint32(y), uint32(w), uint32(h)))
  const M = C.ulong(1 << 32 - 1)
  ximg := C.XGetImage (dpy, C.Drawable(X.win), C.int(x), C.int(y), C.uint(w), C.uint(h), M, C.XYPixmap)
  var pixel C.ulong
  for y := 0; y < int(h); y++ {
    for x := 0; x < int(w); x++ {
      pixel = C.xGetPixel (ximg, C.int(x), C.int(y))
      copy (bs[n:n+4], obj.Encode (int32(pixel)))
      n += 4
    }
  }
// C.xDestroyImage (ximg) // TODO
  return bs
}

func (X *xwindow) Decode (bs []byte) {
  if bs == nil { return }
  n := int32(4 * 4)
  x, y, w, h := obj.Decode4 (bs[:n])
  const M = C.ulong(1 << 32 - 1)
//////////////////////////////////////////////////////////////////////////////////////////////////////////  steals a lot of time
  ximg := C.XGetImage (dpy, C.Drawable(X.win), C.int(x), C.int(y), C.uint(w), C.uint(h), M, C.XYPixmap)
//////////////////////////////////////////////////////////////////////////////////////////////////////////  steals a lot of time
  var pixel C.ulong
  for j := uint32(0); j < uint32(h); j++ {
    for i := uint32(0); i < uint32(w); i++ {
      pixel = (C.ulong)(obj.Decode (n, bs[n:n+4]).(int32))
      C.xPutPixel (ximg, C.int(i), C.int(j), pixel)
      n += 4
    }
  }
  C.XPutImage (dpy, C.Drawable(X.win), X.gc, ximg, 0, 0, C.int(x), C.int(y), C.uint(w), C.uint(h))
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.buffer), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h), C.int(x), C.int(y))
  C.xDestroyImage (ximg)
  C.XFlush (dpy)
}
