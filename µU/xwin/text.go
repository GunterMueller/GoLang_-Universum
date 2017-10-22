package xwin

// (c) Christian Maurer   v. 170814 - license see µU.go

// #include <stdlib.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
import
  "C"
import (
  "unsafe"
  "µU/z"
)

func (X *xwindow) Write1 (b byte, l, c uint) {
  X.Write (z.String(b), l, c)
}

func (X *xwindow) Write (s string, l, c uint) {
  X.WriteGr (s, int(c) * int(X.wd1), int(l) * int(X.ht1))
}

func (X *xwindow) WriteNat (n uint, l, c uint) {
  t := "00"
  if n > 0 {
    const M = 10
    bs := make ([]byte, M)
    for i := 0; i < M; i++ {
      bs[M - 1 - i] = byte('0' + n % 10)
      n = n / M
    }
    s := 0
    for s < M && bs[s] == '0' {
      s++
    }
    t = ""
    if s == M - 1 { s = M - 2 }
    for i := s; i < M - int(n); i++ {
      t += string(bs[i])
    }
  }
  X.Write (t, l, c)
}

func (X *xwindow) Write1Gr (b byte, x, y int) {
  X.WriteGr (z.String(b), x, y)
}

func (X *xwindow) WriteGr (s string, x, y int) {
  C.XSetFont (dpy, X.gc, C.Font(X.fsp.fid)) // TODO TODO TODO
  n := C.uint(len (s))
  if ! X.transparent {
    C.XSetForeground (dpy, X.gc, cc (X.cB))
    if ! X.buff { C.XFillRectangle (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y), n * X.wd1, X.ht1) }
    C.XFillRectangle (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y), n * X.wd1, X.ht1)
    C.XSetForeground (dpy, X.gc, cc (X.cF))
  }
  cs := C.CString (s); defer C.free (unsafe.Pointer (cs))
  if ! X.buff { C.XDrawString (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y) + C.int(X.bl1), cs, C.int(n)) }
  C.XDrawString (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y) + C.int(X.bl1), cs, C.int(n))
  C.XFlush (dpy)
}

func (X *xwindow) Write1InvGr (b byte, x, y int) {
  X.WriteInvGr (z.String(b), x, y)
}

func (X *xwindow) WriteInvGr (s string, x, y int) {
  C.XSetFunction (dpy, X.gc, C.GXinvert)
  X.WriteGr (s, x, y)
  C.XSetFunction (dpy, X.gc, C.GXcopy)
}

func (X *xwindow) Transparent() bool {
  return X.transparent
}

func (X *xwindow) Transparence (t bool) {
  X.transparent = t
}
