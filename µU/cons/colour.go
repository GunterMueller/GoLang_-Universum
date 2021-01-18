package cons

// (c) Christian Maurer   v. 210106 - license see µU.go

import
  "µU/col"
const
  colC = esc1 + "3%d;4%d"
var
  colourdepth uint
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
  c := b.EncodeInv()
  a := uint(0)
  for x := uint(0); x < width; x++ {
    for y := uint(0); y < height; y++ {
      copy (emptyBackground[a:a+colourdepth], c)
      a += colourdepth
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
  X.cF, X.codeF = f, f.EncodeInv()
  X.cB, X.codeB = b, b.EncodeInv()
}

func (X *console) ColourF (f col.Colour) {
  X.cF, X.codeF = f, f.EncodeInv()
}

func (X *console) ColourB (b col.Colour) {
  X.cB, X.codeB = b, b.EncodeInv()
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
  x += uint(X.x)
  y += uint(X.y)
  i := int(width * y + x) * int(colourdepth)
  s := fbcop [i:i+int(colourdepth)]
  c := col.New()
  c.Set (s[0], s[1], s[2])
  return c
}
