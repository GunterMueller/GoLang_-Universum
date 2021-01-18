package scr

// (c) Christian Maurer   v. 210106 - license see µU.go

import (
  "sync"
  . "µU/linewd"
  "µU/mode"
  . "µU/shape"
  "µU/font"
  "µU/col"
  "µU/xwin"
  "µU/cons"
)
type
  screen struct {
                xwin.XWindow
                cons.Console
    colourdepth uint
                mode.Mode
           x, y int
         wd, ht uint
    transparent bool
     blinkMutex sync.Mutex
 blinkX, blinkY uint
    cursorShape,
     blinkShape Shape
       blinking bool
      ppmheader string
             lh uint // len of ppmheader
                }
var (
  underX = xwin.UnderX()
  actual *screen
  modeMax mode.Mode
  width, height = uint(0), uint(0)
)

func newMax() Screen {
  return new_(0, 0, mode.ModeOf (MaxRes()))
}

func (X *screen) init_(w, h, x, y uint) {
  actual = X
  X.wd, X.ht = w, h
  X.x, X.y = int(x), int(y)
  wm, _ := maxRes()
  if wm >= 3840 { // mode.UHD
    X.SetFontsize (font.Huge)
  } else {
    X.SetFontsize (font.Normal)
  }
  X.SetLinewidth (Thin)
  X.Transparence (false)
  X.Colours (col.StartCols())
  X.ScrColours (col.StartCols())
  X.Cls()
}

func new_(x, y uint, m mode.Mode) Screen {
  X := new(screen)
  X.Mode = m
  w, h := mode.Wd(m), mode.Ht(m)
  if underX {
    X.XWindow = xwin.New (x, y, m)
    X.colourdepth = 4 // xwin.ColourDepth()
    width, height = xwin.MaxRes()
  } else {
    X.Console = cons.New (x, y, m)
    X.colourdepth = cons.ColourDepth()
  X.colourdepth = 3 // XXX
    width, height = cons.MaxRes()
  }
  X.init_(w, h, x, y)
  return X
}

func newWH (x, y, w, h uint) Screen {
  X := new(screen)
  X.Mode = mode.None
  if underX {
    X.XWindow = xwin.NewWH (x, y, w, h)
    width, height = w, h
  } else {
    X.Console = cons.NewWH (x, y, w, h)
    width, height = cons.MaxRes()
  }
  X.init_(w, h, x, y)
  return X
}

func (x *screen) Name (n string) {
  if underX {
    x.XWindow.Name (n)
  }
}

func (x *screen) Flush() {
  if underX {
    x.XWindow.Flush()
  }
}

func (x *screen) Fin() {
  if underX {
//  TODO x.Window.Fin()
  } else {
    x.Console.Fin()
  }
}

func maxRes() (uint, uint) {
  var w, h uint
  if underX {
    w, h = xwin.MaxRes()
  } else {
    w, h = cons.MaxRes()
  }
  return w, h
}

func ok (m mode.Mode) bool {
  w, h := maxRes()
  return mode.Wd (m) <= w && mode.Ht(m) <= h
}
