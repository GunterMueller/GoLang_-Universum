package xwin

// (c) murus.org  v. 170816 - license see murus.go

// #include <GL/glx.h>
import
  "C"

func (X *xwindow) WriteGlx() {
  C.glXSwapBuffers (dpy, C.GLXDrawable(X.win))
  X.win2buf()
}
