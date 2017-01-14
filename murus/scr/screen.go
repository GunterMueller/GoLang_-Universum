package scr

// (c) murus.org  v. 140615 - license see murus.go

import (
  "strconv"; "sync"
  "murus/ker"
  . "murus/linewd"; "murus/mode"; . "murus/shape"
  "murus/col"; "murus/font"
  "murus/xker"; "murus/cons"
)
type
  screen struct {
                xker.Window
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
  underX = xker.UnderX()
  actual *screen
  scrList []*screen
  modeMax mode.Mode
  width, height = uint(0), uint(0)
)

func newScr (x, y uint, m mode.Mode) Screen {
  X:= new(screen)
  actual = X
  scrList = append (scrList, X)
  X.Mode = m
  X.x, X.y = int(x), int(y)
  X.wd, X.ht = mode.Wd(m), mode.Ht(m)
  if underX {
    X.Window = xker.New (x, y, m)
    width, height = xker.MaxRes()
  } else {
    X.Console = cons.New (x, y, m)
    width, height = cons.MaxRes()
  }
//  X.Colours (col.White, col.Black)
  X.Colours (col.StartCols())
//  X.ScrColours (col.White, col.Black)
  X.ScrColours (col.StartCols())
//  X.Cls()
  X.SetFontsize (font.Normal)
  X.SetLinewidth (Thin)
  X.Transparence (false)
  return X
}

func newMax() Screen {
  return newScr (0, 0, mode.Mode(maxMode()))
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
    x.Window.Name (n)
  } else {
// TODO
  }
}

func (x *screen) Flush() {
  if underX {
    x.Window.Flush()
  } else {
    // stellt sich das Problem nicht
  }
}

func n() uint {
  return uint(len (scrList))
}

func act() Screen {
  n:= len (scrList)
  a:= n
  if underX {
    for i, s:= range scrList {
      if s.Window == xker.Act() {
        a = i
        break
      }
    }
    if a == n { ker.Panic ("no actual screen found") }
  } else {
    a = cons.ActIndex()
  }
  actual = scrList[a]
  for i:= a; i < n - 1; i++ {
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
    width, height = xker.MaxRes()
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
  for m:= mode.Mode(0); m < mode.NModes; m++ {
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

func (X *screen) Move (x, y int) {
  if underX { return } // da stellt sich das Problem nicht
/*
  ok:= 0 <= X.x + x                          && 0 <= X.y + y                           &&
            X.x + x + int(X.wd) < int(width) &&      X.y + y + int(X.ht) < int(height)

  ok:= 0 <= x                          && 0 <= y                           &&
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
