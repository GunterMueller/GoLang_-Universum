package cons

// (c) murus.org  v. 140217 - license see murus.go

import
  "murus/font"

func (X *console) ActFontsize() font.Size {
  return X.fontsize
}

func (X *console) SetFontsize (f font.Size) {
  X.fontsize = f
  X.ht1, X.wd1 = font.Ht (X.fontsize), font.Wd (X.fontsize)
  X.nLines, X.nColumns = X.ht / X.ht1, X.wd / X.wd1
}
