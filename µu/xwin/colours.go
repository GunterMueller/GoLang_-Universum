package xwin

// (c) Christian Maurer   v. 170814 - license see µu.go

// #include <X11/X.h>
// #include <X11/Xlib.h>
import
  "C"
import (
  "µu/col"
)

func (X *xwindow) ScrColours (f, b col.Colour) {
  X.scrF, X.scrB = f, b
}

func (X *xwindow) ScrColourF (f col.Colour) {
  X.scrF = f
}

func (X *xwindow) ScrColourB (b col.Colour) {
  X.scrB = b
}

func (X *xwindow) ScrCols() (col.Colour, col.Colour) {
  return X.scrF, X.scrB
}

func (X *xwindow) ScrColF() col.Colour {
  return X.scrF
}

func (X *xwindow) ScrColB() col.Colour {
  return X.scrB
}

func (X *xwindow) Colours (f, b col.Colour) {
  if ! initialized { panic ("xwin.Colours: ! initialized"); return }
  X.cF, X.cB = f, b
  C.XSetForeground (dpy, X.gc, cc (X.cF))
  C.XSetBackground (dpy, X.gc, cc (X.cB))
}

func (X *xwindow) ColourF (f col.Colour) {
//  print ("/ ")
  X.cF = f
  C.XSetForeground (dpy, X.gc, cc (X.cF))
//  println ("// ")
}

func (X *xwindow) ColourB (b col.Colour) {
  X.cB = b
  C.XSetBackground (dpy, X.gc, cc (X.cB))
}

func (X *xwindow) Cols() (col.Colour, col.Colour) {
  return X.cF, X.cB
}

func (X *xwindow) ColF() col.Colour {
  return X.cF
}

func (X *xwindow) ColB() col.Colour {
  return X.cB
}

func (X *xwindow) Colour (x, y uint) col.Colour {
  return X.scrB
}
