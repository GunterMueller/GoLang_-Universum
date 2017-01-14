package scr

// (c) murus.org  v. 140527 - license see murus.go

func (X *screen) WriteGlx() {
  if underX {
    X.Window.WriteGlx()
  } else {
    X.Console.WriteGlx()
  }
}
