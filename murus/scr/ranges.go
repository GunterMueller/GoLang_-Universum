package scr

// (c) murus.org  v. 140527 - license see murus.go

func (X *screen) Cls() {
  if underX {
    X.Window.Cls()
  } else {
    X.Console.Cls()
  }
}

func (X *screen) Clr (l, c, w, h uint) {
  if underX {
    X.Window.Clr (l, c, w, h)
  } else {
    X.Console.Clr (l, c, w, h)
  }
}

func (X *screen) ClrGr (x, y, x1, y1 int) {
  if underX {
    X.Window.ClrGr (x, y, x1, y1)
  } else {
    X.Console.ClrGr (x, y, x1, y1)
  }
}

func (x *screen) Buf (on bool) {
  if underX {
    x.Window.Buf (on)
  } else {
    x.Console.Buf (on)
  }
}

func (x *screen) Buffered () bool {
  if underX {
    return x.Window.Buffered()
  } else {
    return x.Console.Buffered()
  }
}

func (X *screen) Save (l, c, w, h uint) {
  if underX {
    X.Window.Save (l, c, w, h)
  } else {
    X.Console.Save (l, c, w, h)
  }
}

func (X *screen) SaveGr (x, y, x1, y1 int) {
  if underX {
    X.Window.SaveGr (x, y, x1, y1)
  } else {
    X.Console.SaveGr (x, y, x1, y1)
  }
}

func (X *screen) Save1() {
  if underX {
    X.Window.Save1()
  } else {
    X.Console.Save1()
  }
}

func (X *screen) Restore (l, c, w, h uint) {
  if underX {
    X.Window.Restore (l, c, w, h)
  } else {
    X.Console.Restore (l, c, w, h)
  }
}

func (X *screen) RestoreGr (x, y, x1, y1 int) {
  if underX {
    X.Window.RestoreGr (x, y, x1, y1)
  } else {
    X.Console.RestoreGr (x, y, x1, y1)
  }
}

func (X *screen) Restore1() {
  if underX {
    X.Window.Restore1()
  } else {
    X.Console.Restore1()
  }
}
