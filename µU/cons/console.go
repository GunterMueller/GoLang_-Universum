package cons

// (c) Christian Maurer   v. 210103 - license see µU.go

import (
  "sync"
  "strconv"
  . "µU/obj"
  . "µU/linewd"
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
         archive Stream
          shadow []Stream
            buff bool
        wd1, ht1 uint
cF, cB, cFA, cBA col.Colour
    codeF, codeB Stream
      scrF, scrB col.Colour
          lineWd Linewidth
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
         polygon [][]bool // to fill polygons
            done [][]bool // to fill polygons
        incident bool
   xx_, yy_, tt_ int // for incidence tests
                 }
var (
  actMutex sync.Mutex
  actual *console
  mouseIndex int
  mouseConsole *console
  width, height uint
)

func (X *console) ok() bool {
  return uint(X.x) + X.wd <= width && uint(X.y) + X.ht <= height
}

func (X *console) init_(x, y uint) {
  actual = X
  mouseConsole = X
  X.x, X.y = int(x), int(y)
  X.cF, X.cB = col.StartCols()
  X.cFA, X.cBA = col.StartColsA()
  if ! framebufferOk() {
    ker.Panic ("µU does not support far tty-consoles")
  }
  if ! X.ok() { a, b, c := strconv.Itoa(X.x), strconv.Itoa(int(X.wd)), strconv.Itoa (int(width))
                d, e, f := strconv.Itoa(X.y), strconv.Itoa(int(X.ht)), strconv.Itoa (int(height))
    ker.Panic ("new console too large: " + a + " + " + b + " > " + c + " or " +
                                           d + " + " + e + " > " + f)
  }
  X.archive = make(Stream, fbmemsize)
  X.shadow = make([]Stream, X.ht)
  for i := 0; i < int(X.ht); i++ {
    X.shadow[i] = make(Stream, X.wd * colourdepth)
  }
  X.initMouse()
  X.polygon = make([][]bool, X.wd)
  X.done = make([][]bool, X.wd)
  for i := 0; i < int(X.wd); i++ {
    X.polygon[i] = make([]bool, X.ht)
    X.done[i] = make([]bool, X.ht)
  }
  X.ScrColours (X.cF, X.cB)
  X.Cls()
  X.SetFontsize (font.Normal)
  X.doBlink()
}

func new_(x, y uint, m Mode) Console {
  X := new(console)
  X.wd, X.ht = Wd(m), Ht(m)
  X.init_(x, y)
  return X
}

func newWH (x, y, w, h uint) Console {
  X := new(console)
  X.wd, X.ht = w, h
  X.init_(x, y)
  return X
}

func newMax() Console {
  return new_(0, 0, ModeOf(maxRes()))
}
