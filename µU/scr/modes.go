package scr

// (c) Christian Maurer   v. 170818 - license see µU.go

import
  "µU/mode"

func (X *screen) ActMode() mode.Mode {
  return X.Mode
}

func (x *screen) X() uint {
  return uint(x.x)
}

func (x *screen) Y() uint {
  return uint(x.y)
}

func (x *screen) Wd() uint{
  return x.wd
}

func (x *screen) Ht() uint{
  return x.ht
}

func (x *screen) Wd1() uint{
  if underX {
    return x.XWindow.Wd1()
  }
  return x.Console.Wd1()
}

func (x *screen) Ht1() uint{
  if underX {
    return x.XWindow.Ht1()
  }
  return x.Console.Ht1()
}

func (x *screen) NLines() uint{
  if underX {
    return x.XWindow.NLines()
  }
  return x.Console.NLines()
}

func (x *screen) NColumns() uint{
  if underX {
    return x.XWindow.NColumns()
  }
  return x.Console.NColumns()
}

func (x *screen) Proportion() float64 {
  return float64 (x.wd) / float64 (x.ht)
}
