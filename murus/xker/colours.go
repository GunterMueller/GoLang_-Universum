package xker

// (c) murus.org  v. 140217 - license see murus.go

// #include <X11/X.h>
// #include <X11/Xlib.h>
import
  "C"
import (
  "murus/col"
)

func (X *window) ScrColours (f, b col.Colour) {
  X.scrF, X.scrB = f, b
}

func (X *window) ScrColourF (f col.Colour) {
  X.scrF = f
}

func (X *window) ScrColourB (b col.Colour) {
  X.scrB = b
}

func (X *window) ScrCols() (col.Colour, col.Colour) {
  return X.scrF, X.scrB
}

func (X *window) ScrColF() col.Colour {
  return X.scrF
}

func (X *window) ScrColB() col.Colour {
  return X.scrB
}

func (X *window) Colours (f, b col.Colour) {
  if ! initialized { panic ("xker.Colours: ! initialized"); return }
  X.cF, X.cB = f, b
  C.XSetForeground (dpy, X.gc, cc (X.cF))
  C.XSetBackground (dpy, X.gc, cc (X.cB))
}

func (X *window) ColourF (f col.Colour) {
//  print ("/ ")
  X.cF = f
  C.XSetForeground (dpy, X.gc, cc (X.cF))
//  println ("// ")
}

func (X *window) ColourB (b col.Colour) {
  X.cB = b
  C.XSetBackground (dpy, X.gc, cc (X.cB))
}

func (X *window) Cols() (col.Colour, col.Colour) {
  return X.cF, X.cB
}

func (X *window) ColF() col.Colour {
  return X.cF
}

func (X *window) ColB() col.Colour {
  return X.cB
}

func (X *window) Colour (x, y uint) col.Colour {
  return X.scrB
}
