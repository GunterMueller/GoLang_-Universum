package scr

// (c) Christian Maurer   v. 170814 - license see µU.go

import
  . "µU/ptr"

func (X *screen) MouseEx() bool {
  if underX {
    return X.XWindow.MouseEx()
  }
  return X.Console.MouseEx()
}

func (X *screen) SetPointer (p Pointer) {
  if underX {
    X.XWindow.SetPointer (p)
  } else {
    X.Console.SetPointer (p)
  }
}

func (X *screen) MousePos() (uint, uint) {
  if underX {
    return X.XWindow.MousePos()
  }
  return X.Console.MousePos()
}

func (X *screen) MousePosGr() (int, int) {
  if underX {
    return X.XWindow.MousePosGr()
  }
  return X.Console.MousePosGr()
}

func (X *screen) WarpMouse (l, c uint) {
  if underX {
    X.XWindow.WarpMouse (l, c)
  } else {
    X.Console.WarpMouse (l, c)
  }
}

func (X *screen) WarpMouseGr (x, y int) {
  if underX {
    X.XWindow.WarpMouseGr (x, y)
  } else {
    X.Console.WarpMouseGr (x, y)
  }
}

func (X *screen) MousePointer (on bool) {
  if underX {
    X.XWindow.MousePointer (on)
  } else {
    X.Console.MousePointer (on)
  }
}

func (X *screen) MousePointerOn() bool {
  if underX {
    return X.XWindow.MousePointerOn()
  }
  return X.Console.MousePointerOn()
}

func (X *screen) UnderMouse (l, c, w, h uint) bool {
  if underX {
    return X.XWindow.UnderMouse (l, c, w, h)
  }
  return X.Console.UnderMouse (l, c, w, h)
}

func (X *screen) UnderMouseGr (x, y, x1, y1 int, t uint) bool {
  if underX {
    return X.XWindow.UnderMouseGr (x, y, x1, y1, t)
  }
  return X.Console.UnderMouseGr (x, y, x1, y1, t)
}

func (X *screen) UnderMouse1 (x, y int, r uint) bool {
  if underX {
    return X.XWindow.UnderMouse1 (x, y, r)
  }
  return X.Console.UnderMouse1 (x, y, r)
}
