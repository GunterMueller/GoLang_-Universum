package scr

// (c) Christian Maurer   v. 170917 - license see µu.go

import (
  "strconv"
  "sync"
  "µu/ker"
  . "µu/linewd"
  "µu/mode"
  . "µu/shape"
  "µu/font"
  "µu/col"
  "µu/xwin"
  "µu/cons"
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
                }
var (
  underX = xwin.UnderX()
  actual *screen
  scrList []*screen
  modeMax mode.Mode
  width, height = uint(0), uint(0)
)

func newMax() Screen {
  return newScr (0, 0, mode.Mode(maxMode()))
}

func newScr (x, y uint, m mode.Mode) Screen {
  if m == mode.WH { ker.Panic ("use newWH !") }
  X := new(screen)
  actual = X
  scrList = append (scrList, X)
  X.Mode = m
  X.x, X.y = int(x), int(y)
  X.wd, X.ht = mode.Wd(m), mode.Ht(m)
  if underX {
    X.XWindow = xwin.New (x, y, m)
    width, height = xwin.MaxRes()
  } else {
    X.Console = cons.New (x, y, m)
    width, height = cons.MaxRes()
  }
  X.Colours (col.StartCols())
  X.ScrColours (col.StartCols())
  X.Cls()

  X.SetFontsize (font.Normal)
  X.SetLinewidth (Thin)
  X.Transparence (false)
  return X
}

func newWH (x, y, w, h uint) Screen {
  X := new(screen)
  actual = X
  scrList = append (scrList, X)
  X.Mode = mode.WH
  mode.WdHt (w, h)
  X.x, X.y = int(x), int(y)
  X.wd, X.ht = w, h
  if underX {
    X.XWindow = xwin.New (x, y, mode.WH)
    width, height = xwin.MaxRes()
  } else {
//    X.Console = cons.NewWH (x, y, w, h)
//    width, height = cons.MaxRes()
    ker.Panic ("newWH is not yet implemented for tty-concoles")
  }
  X.Colours (col.StartCols())
  X.ScrColours (col.StartCols())
  X.Cls()
  X.SetFontsize (font.Normal)
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

// experimental ////////////////////////////////////////////////////////

func (X *screen) Movex (x, y int) {
  if underX { return } // da stellt sich das Problem nicht
/*
  ok := 0 <= X.x + x                          && 0 <= X.y + y                           &&
             X.x + x + int(X.wd) < int(width) &&      X.y + y + int(X.ht) < int(height)

  ok := 0 <= x                          && 0 <= y                           &&
             x + int(X.wd) < int(width) &&      y + int(X.ht) < int(height)

  if ! ok { return }
*/
  if X.Console.Moved (x, y) {
    X.x += x
    X.y += y
/*
    if X.x != int(X.X()) { ker.Panic ("Katzenkacke") }
    if X.y != int(X.Y()) { ker.Panic ("Schweinescheiße") }
*/
  }
}

func init() {
//  var _ Window = newMax()
//  visible = true // TODO -> only in initConsole()
}

func start (x, y, z, xf, yf, zf float64) {
  if underX {
//    gl.Name()
    actual.Start (x, y, z, xf, yf, zf)
  } else {
    ker.Panic ("no GUI")
  }
}

func move (d int, a float64) {
  if underX {
    actual.Move (d, a)
  } else {
    ker.Panic ("no GUI")
  }
}

func turnAroundFocus (d int, a float64) {
  if underX {
    actual.TurnAroundFocus (d, a)
  } else {
    ker.Panic ("no GUI")
  }
}

func draw (d func()) {
  if underX {
    actual.Draw (d)
  } else {
    ker.Panic ("no GUI")
  }
}

func look (d func()) {
  if underX {
    actual.Look (d)
  } else {
    ker.Panic ("no GUI")
  }
}
