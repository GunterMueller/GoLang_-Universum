package xwin

// (c) Christian Maurer   v. 201016 - license see µU.go

// #include <stdio.h>
// #include <stdlib.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/Xlibint.h>
/*
void copy (Display *d, char* s, int n) {
  XStoreBytes (d, s, n);
}
char *paste (Display *d, int* n) {
  return XFetchBytes (d, n);
}
void xfree (char* s) {
  Xfree (s);
}
*/
import
  "C"
import
  "unsafe"

func (X *xwindow) Copy (s string) {
  cs, n := C.CString (s), C.int(len (s))
  defer C.free (unsafe.Pointer (cs))
  C.copy (dpy, cs, n)
}

func (X *xwindow) Paste() string {
  var (cs *C.char; n C.int)
  defer C.free (unsafe.Pointer (cs))
  cs = C.paste (dpy, &n)
  s := C.GoStringN (cs, n)
  C.xfree (cs)
  X.Flush()
  return s
}
