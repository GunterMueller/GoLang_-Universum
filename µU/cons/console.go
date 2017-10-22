package cons

// (c) Christian Maurer   v. 170810 - license see µU.go

import (
  "sync"
  "strconv"
//  . "µU/linewd"
  . "µU/shape"
  . "µU/mode"
  "µU/ker"
  "µU/col"
  "µU/font"
  "µU/ptr"
)
type
  console struct {
            x, y int
          wd, ht uint
nLines, nColumns uint
         archive []byte
          shadow [][]byte
            buff bool
        wd1, ht1 uint
cF, cB, cFA, cBA col.Colour
    codeF, codeB []byte
      scrF, scrB col.Colour
//          lineWd Linewidth
        fontsize font.Size
     transparent bool
     cursorShape,
    consoleShape,
      blinkShape Shape
  blinkX, blinkY uint
      blinkMutex sync.Mutex
        blinking bool
         mouseOn bool
         pointer ptr.Pointer
  xMouse, yMouse int
              pg [][]bool // to fill polygons
        incident bool
   xx_, yy_, tt_ int // for incidence tests
             }
var (
  actMutex sync.Mutex
  consList []*console
  actual *console
  mouseIndex int
  mouseConsole *console
  first = true
  width, height uint
)

func imp (x, y uint) *console {
  for _, s := range consList {
    if s.x <= int(x) && x <= s.wd && s.y <= int(y) && y <= s.ht {
      return s
    }
  }
  ker.Panic ("µU/scr/imp: there is no console there")
  return nil
}

func (X *console) ok() bool {
  return uint(X.x) + X.wd <= width && uint(X.y) + X.ht <= height
}

func newCons (x, y uint, m Mode) Console {
  X := new(console)
  actual = X
  mouseConsole = X
  X.x, X.y = int(x), int(y)
  X.wd, X.ht = Wd(m), Ht(m)
  X.cF, X.cB = col.StartCols()
  X.cFA, X.cBA = col.StartColsA()
  consoleInit()
  if ! X.ok() {
    ker.Panic ("new console too large: " +
               strconv.Itoa(X.x) + " + " + strconv.Itoa(int(X.wd)) + " > " + strconv.Itoa (int(width)) + " or " +
               strconv.Itoa(X.y) + " + " + strconv.Itoa(int(X.ht)) + " > " + strconv.Itoa (int(height)))
  }
  X.archive = make ([]byte, fbmemsize)
  X.shadow = make ([][]byte, X.ht)
  for i := 0; i < int(X.ht); i++ { X.shadow[i] = make ([]byte, X.wd * colourdepth) }
  X.initMouse()
  X.pg = make ([][]bool, X.ht)
  for i := 0; i < int(X.ht); i++ { X.pg[i] = make ([]bool, X.wd) }
  X.ScrColours (X.cF, X.cB)
  X.Cls()
  if first { defer goMouse(); first = false }
  X.SetFontsize (font.Normal)
  X.doBlink()
  actMutex.Lock()
  consList = append (consList, X)
  actMutex.Unlock()
  return X
}

func newMax() Console {
  return newCons (0, 0, ModeOf(maxRes()))
}

func goMouse() {
  go followMouse()
}

func actIndex() int {
  actMutex.Lock()
  defer actMutex.Unlock()
  n := len (consList)
  for i := mouseIndex; i < n - 1; i++ {
    consList[i] = consList[i + 1]
  }
  consList[n - 1] = mouseConsole
  actual = consList[n - 1]
  actual.MousePointer (true)
  return mouseIndex
}

func followMouse() {
  for {
    i := len(consList) - 1
    actMutex.Lock()
    for {
      s := consList[i]
      xm, ym := s.MousePosGr()
      if 0 <= xm && xm <= int(s.wd) && 0 <= ym && ym <= int(s.ht) {
        mouseIndex, mouseConsole = i, s
        break // if consoles overlap, the last (= newest) one is the actual one
      }
      if i == 0 {
        break
      }
      i --
    }
    actMutex.Unlock()
    ker.Msleep (100)
  }
}
