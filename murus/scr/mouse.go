package scr

// (c) murus.org  v. 160109 - license see murus.go

import
  . "murus/ptr"

func (X *screen) MouseEx() bool {
  if underX {
    return X.Window.MouseEx()
  }
  return X.Console.MouseEx()
}

func (X *screen) SetPointer (p Pointer) {
  if underX {
    X.Window.SetPointer (p)
  } else {
    X.Console.SetPointer (p)
  }
}

func (X *screen) MousePos() (uint, uint) {
  if underX {
    return X.Window.MousePos()
  }
  return X.Console.MousePos()
}

func (X *screen) MousePosGr() (int, int) {
  if underX {
    return X.Window.MousePosGr()
  }
  return X.Console.MousePosGr()
}

func (X *screen) WarpMouse (l, c uint) {
  if underX {
    X.Window.WarpMouse (l, c)
  } else {
    X.Console.WarpMouse (l, c)
  }
}

func (X *screen) WarpMouseGr (x, y int) {
  if underX {
    X.Window.WarpMouseGr (x, y)
  } else {
    X.Console.WarpMouseGr (x, y)
  }
}

func (X *screen) MousePointer (on bool) {
  if underX {
    X.Window.MousePointer (on)
  } else {
    X.Console.MousePointer (on)
  }
}

func (X *screen) MousePointerOn() bool {
  if underX {
    return X.Window.MousePointerOn()
  }
  return X.Console.MousePointerOn()
}

func (X *screen) UnderMouse (l, c, w, h uint) bool {
  if underX {
    return X.Window.UnderMouse (l, c, w, h)
  }
  return X.Console.UnderMouse (l, c, w, h)
}

func (X *screen) UnderMouseGr (x, y, x1, y1 int, t uint) bool {
  if underX {
    return X.Window.UnderMouseGr (x, y, x1, y1, t)
  }
  return X.Console.UnderMouseGr (x, y, x1, y1, t)
}

func (X *screen) UnderMouse1 (x, y int, r uint) bool {
  if underX {
    return X.Window.UnderMouse1 (x, y, r)
  }
  return X.Console.UnderMouse1 (x, y, r)
}
