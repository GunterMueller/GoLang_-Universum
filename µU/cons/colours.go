package cons

// (c) Christian Maurer   v. 170918 - license see µU.go

import
  "µU/col"
const
  colC = esc1 + "3%d;4%d"
var
  colourdepth uint // 1..4 in Byte
/*
func startCols() (col.Colour, col.Colour) {
  return col.White, col.Black
}

func startColsA() (col.Colour, col.Colour) {
  return col.Red, col.Black
}
*/
func (X *console) ScrColours (f, b col.Colour) {
  X.scrF = f
  X.ScrColourB (b)
}

func (X *console) ScrColourF (f col.Colour) {
  X.scrF = f
}

func (X *console) ScrColourB (b col.Colour) {
  X.scrB = b.Clone().(col.Colour)
  c := b.Cc()
  a := 0
  for x := 0; x < int(width); x++ {
    for y := 0; y < int(height); y++ {
      copy (emptyBackground[a:a+int(colourdepth)], c)
      a += int(colourdepth)
    }
  }
}

func (X *console) ScrCols() (col.Colour, col.Colour) {
  return X.scrF, X.scrB
}

func (X *console) StartCols() (col.Colour, col.Colour) {
  return X.cF, X.cB
}

func (X *console) StartColsA() (col.Colour, col.Colour) {
  return X.cFA, X.cBA
}

func (X *console) ScrColF() col.Colour {
  return X.scrF
}

func (X *console) ScrColB() col.Colour {
  return X.scrB
}

func (X *console) Colours (f, b col.Colour) {
  X.cF, X.codeF = f, f.Cc()
  X.cB, X.codeB = b, b.Cc()
}

func (X *console) ColourF (f col.Colour) {
  X.cF, X.codeF = f, f.Cc()
}

func (X *console) ColourB (b col.Colour) {
  X.cB, X.codeB = b, b.Cc()
}

func (X *console) Cols() (col.Colour, col.Colour) {
  return X.cF, X.cB
}

func (X *console) ColF() col.Colour {
  return X.cF
}

func (X *console) ColB() col.Colour {
  return X.cB
}

func (X *console) Colour (x, y uint) col.Colour {
  if x >= X.wd || y >= X.ht || ! visible {
    return X.scrB
  }
  x += uint(X.x); y += uint(X.y)
  a := int(width * y + x) * int(colourdepth)
  return col.P6Colour (fbcop [a:a+int(colourdepth)])
}
