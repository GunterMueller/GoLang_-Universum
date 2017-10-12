package cons

// (c) Christian Maurer   v. 170818 - license see µu.go

import
  "µu/font"

func (X *console) ActFontsize() font.Size {
  return X.fontsize
}

func (X *console) SetFontsize (f font.Size) {
  X.fontsize = f
  X.ht1, X.wd1 = font.Ht (X.fontsize), font.Wd (X.fontsize)
  X.nLines, X.nColumns = X.ht / X.ht1, X.wd / X.wd1
}

func (x *console) Wd1() uint{
  return x.wd1
}

func (x *console) Ht1() uint{
  return x.ht1
}

func (x *console) NLines() uint{
  return x.nLines
}

func (x *console) NColumns() uint{
  return x.nColumns
}
