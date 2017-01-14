package xker

// (c) murus.org  v. 140217 - license see murus.go

// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <GL/gl.h>
// #include <GL/glx.h>
/*
void write (Display *d, Window w) { glXSwapBuffers (d, w); }
*/
import
  "C"

func (x *window) WriteGlx() {
  C.write (dpy, x.win)
  x.win2buf()
}
