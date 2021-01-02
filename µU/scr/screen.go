package scr

// (c) Christian Maurer   v. 201031 - license see µU.go

import (
  "strconv"
  "sync"
  "µU/ker"
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
                mode.Mode
           x, y int
         wd, ht uint
    transparent bool
     blinkMutex sync.Mutex
 blinkX, blinkY uint
    cursorShape,
     blinkShape Shape
       blinking bool
//        phantom [][]col.Colour
                }
var (
  underX = xwin.UnderX()
  actual *screen
  scrList []*screen
  modeMax mode.Mode
  width, height = uint(0), uint(0)
)

func newMax() Screen {
  return new_(0, 0, mode.Mode(maxMode()))
}

func new_(x, y uint, m mode.Mode) Screen {
//  if m == mode.WH { ker.Panic ("use newWH !") }
  X := new(screen)
  actual = X
  scrList = append (scrList, X)
  X.Mode = m
  X.x, X.y = int(x), int(y)
  X.wd, X.ht = mode.Wd(m), mode.Ht(m)
////  X.phantom = make([][]col.Colour, X.ht)
//  X.phantom = make([][]col.Colour, 1600)
////  for y := uint(0); y < X.ht; y++ {
//  for y := uint(0); y < 1600; y++ {
////    X.phantom[y] = make([]col.Colour, X.wd)
//    X.phantom[y] = make([]col.Colour, 1600)
//  }
  if underX {
    X.XWindow = xwin.New (x, y, m)
    width, height = xwin.MaxRes()
  } else {
    X.Console = cons.New (x, y, m)
    if X.Console == nil {
      panic ("µU does not yet work on far tty-consoles")
    }
    width, height = cons.MaxRes()
  }
  X.Colours (col.StartCols())
  X.ScrColours (col.StartCols())
  X.Cls()
  if maxMode() == mode.UHD {
    X.SetFontsize (font.Huge)
  } else {
    X.SetFontsize (font.Normal)
  }
  X.SetLinewidth (Thin)
  X.Transparence (false)
  return X
}

func newWH (x, y, w, h uint) Screen {
  X := new(screen)
//  X.phantom = make([][]col.Colour, 1000)
//  for y := uint(0); y < 1000; y++ {
//    X.phantom[y] = make([]col.Colour, 1000)
//  }
  actual = X
  scrList = append (scrList, X)
//  X.Mode = mode.WH
//  mode.WdHt (w, h)
  X.x, X.y = int(x), int(y)
  X.wd, X.ht = w, h
  if underX {
    X.XWindow = xwin.NewWH (x, y, w, h)
    width, height = w, h
  } else {
//    X.Console = cons.NewWH (x/8, y/16, w/8, h/16)
    X.Console = cons.NewWH (x, y, w, h)
    width, height = cons.MaxRes()
  }
  X.Colours (col.StartCols())
  X.ScrColours (col.StartCols())
  X.Cls()
  if maxMode() == mode.UHD {
    X.SetFontsize (font.Huge)
  } else {
    X.SetFontsize (font.Normal)
  }
  X.SetLinewidth (Thin)
  X.Transparence (false)
  return X
}

func (x *screen) fin() {
  if underX {
    // x.Window.Fin() // TODO
  } else {
    x.Console.Fin()
  }
}

func (x *screen) Name (n string) {
  if underX {
    x.XWindow.Name (n)
  } else {
// TODO
  }
}

func (x *screen) Flush() {
  if underX {
    x.XWindow.Flush()
  } else {
    // stellt sich das Problem nicht
  }
}

func n() uint {
  return uint(len (scrList))
}

func act() Screen {
  n := len (scrList)
  a := n
  if underX {
    for i, s := range scrList {
      if s.XWindow == xwin.Act() {
        a = i
        break
      }
    }
    if a == n { ker.Panic ("no actual screen found") }
  } else {
    a = cons.ActIndex()
  }
  actual = scrList[a]
  for i := a; i < n - 1; i++ {
    scrList[i] = scrList[i + 1]
  }
  scrList[n - 1] = actual
  if ! underX {
//    actual.Save1()
  }
  return actual
}

func (x *screen) Fin() {
  if underX {
//  TODO x.Window.Fin()
  } else {
    x.Console.Fin()
  }
}

func maxMode() mode.Mode {
  if underX {
    width, height = xwin.MaxRes()
  } else {
    width, height = cons.MaxRes()
  }
  return modeOf (width, height)
}

func maxRes() (uint, uint) {
  maxMode()
  return width, height
}

func maxX() uint {
  maxMode()
  return width
}

func maxY() uint {
  maxMode()
  return height
}

func modeOf (x, y uint) mode.Mode {
  for m := mode.Mode(0); m < mode.NModes; m++ {
    if mode.Wd(m) == width && mode.Ht(m) == height {
      return m
    }
  }
  ker.Panic ("hardware reports undefined mode " + strconv.Itoa(int(width)) + " x " + strconv.Itoa(int(height)))
  return mode.NModes
}

func ok (m mode.Mode) bool {
  modeMax = maxMode()
  return mode.Wd(m) <= mode.Wd(modeMax) && mode.Ht(m) <= mode.Ht(modeMax)
}

/*/
func (x *screen) Phantom() [][]col.Colour {
  return x.phantom
}

func (X *screen) DrawPhantom() {
//  for y := 0; y < int(X.ht); y++ {
  for y := 0; y < 400; y++ {
//    for x := 0; x < int(X.wd); x++ {
    for x := 0; x < 600; x++ {
      c := X.phantom[x][y]
      ColourF (c)
      Point (x, y)
    }
   }
}
/*/
