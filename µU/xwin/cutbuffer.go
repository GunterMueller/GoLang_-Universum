package xwin

// (c) Christian Maurer   v. 190528 - license see µU.go

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
void copy7 (Display *d, char* s, int n, int b) {
  XStoreBuffer (d, s, n, b);
}
char *paste7 (Display *d, int* n, int b) {
  return XFetchBuffer (d, n, b);
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

func (X *xwindow) Copy7 (s string, i int) {
  if i < 0 || i > 7 { panic("xwin.Copy7") }
  cs, n := C.CString (s), C.int(len (s))
  defer C.free (unsafe.Pointer (cs))
  C.copy7 (dpy, cs, n, C.int(i))
}

func (X *xwindow) Paste7 (i int) string {
  if i < 0 || i > 7 { panic("xwin.Paste7") }
  var (cs *C.char; n C.int)
  defer C.free (unsafe.Pointer (cs))
  cs = C.paste7 (dpy, &n, C.int(i))
  s := C.GoStringN (cs, n)
  C.xfree (cs)
  X.Flush()
  return s
}
