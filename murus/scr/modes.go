package scr

// (c) murus.org  v. 140527 - license see murus.go

import
  "murus/mode"

func (X *screen) ActMode() mode.Mode {
  return X.Mode
/*
  if underX {
    return X.Window.ActMode()
  }
  return X.Console.ActMode()
*/
}

func (x *screen) X() uint {
  return uint(x.y)
  if underX {
    return x.Window.X()
  }
  return x.Console.X()
}

func (x *screen) Y() uint {
  return uint(x.y)
  if underX {
    return x.Window.Y()
  }
  return x.Console.Y()
}

func (x *screen) Wd() uint{
  if underX {
    return x.Window.Wd()
  }
  return x.Console.Wd()
}

func (x *screen) Ht() uint{
  if underX {
    return x.Window.Ht()
  }
  return x.Console.Ht()
}

func (x *screen) Wd1() uint{
  if underX {
    return x.Window.Wd1()
  }
  return x.Console.Wd1()
}

func (x *screen) Ht1() uint{
  if underX {
    return x.Window.Ht1()
  }
  return x.Console.Ht1()
}

func (x *screen) NLines() uint{
  if underX {
    return x.Window.NLines()
  }
  return x.Console.NLines()
}

func (x *screen) NColumns() uint{
  if underX {
    return x.Window.NColumns()
  }
  return x.Console.NColumns()
}

func (x *screen) Proportion() float64 {
  return float64 (x.wd) / float64 (x.ht)
}
