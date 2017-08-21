package scr

// (c) murus.org  v. 170814 - license see murus.go

import
  "murus/col"
var
  colourdepth uint // 1..4 in Byte

func cc (c col.Colour) []byte {
  n:= col.Code (c)
  b:= make([]byte, colourdepth)
  for i:= 0; i < int(colourdepth); i++ {
    b[i] = byte(n)
    n >>= 8
  }
  return b
}

/*
// Pre: len(bs) == colourdepth
func cd (bs[]byte) uint { // inverse of cc
  n:= uint(0)
  if len(bs) == int(colourdepth) {
    for i:= int(colourdepth) - 1; i >= 0; i-- {
      n = n * 1<<8 + uint(bs[i])
    }
  }
  return n
}

func (X *screen) ScrCols() (col.Colour, col.Colour) {
  if underX {
    return X.Window.ScrCols()
  }
  return X.Console.ScrCols()
}
*/

func startCols() (col.Colour, col.Colour) {
  return col.StartCols()
}

func startColsA() (col.Colour, col.Colour) {
  return col.StartColsA()
}

func (X *screen) ScrColours (f, b col.Colour) {
  if underX {
    X.XWindow.ScrColours (f, b)
  } else {
    X.Console.ScrColours (f, b)
  }
}

func (X *screen) ScrColourF (f col.Colour) {
  if underX {
    X.XWindow.ScrColourF (f)
  } else {
    X.Console.ScrColourF (f)
  }
}

func (X *screen) ScrColourB (b col.Colour) {
  if underX {
    X.XWindow.ScrColourB (b)
  } else {
    X.Console.ScrColourB (b)
  }
}

func (X *screen) ScrCols() (col.Colour, col.Colour) {
  if underX {
    return X.XWindow.ScrCols()
  }
  return X.Console.ScrCols()
}

func (X *screen) ScrColF() col.Colour {
  if underX {
    return X.XWindow.ScrColF()
  }
  return X.Console.ScrColF()
}

func (X *screen) ScrColB() col.Colour {
  if underX {
    return X.XWindow.ScrColB()
  }
  return X.Console.ScrColB()
}

func (X *screen) Colours (f, b col.Colour) {
  if underX {
    X.XWindow.Colours (f, b)
  } else {
    X.Console.Colours (f, b)
  }
}

func (X *screen) ColourF (f col.Colour) {
  if underX {
    X.XWindow.ColourF (f)
  } else {
    X.Console.ColourF (f)
  }
}

func (X *screen) ColourB (b col.Colour) {
  if underX {
    X.XWindow.ColourB (b)
  } else {
    X.Console.ColourB (b)
  }
}

func (X *screen) Cols() (col.Colour, col.Colour) {
  if underX {
    return X.XWindow.Cols()
  }
  return X.Console.Cols()
}

func (X *screen) ColF() col.Colour {
  if underX {
    return X.XWindow.ColF()
  }
  return X.Console.ColF()
}

func (X *screen) ColB() col.Colour {
  if underX {
    return X.XWindow.ColB()
  }
  return X.Console.ColB()
}

func (X *screen) Colour (x, y uint) col.Colour {
  if underX {
    return X.XWindow.Colour (x, y)
  }
  return X.Console.Colour (x, y)
}
