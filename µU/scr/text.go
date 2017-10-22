package scr

// (c) Christian Maurer   v. 170814 - license see µU.go

import
  "µU/str"

func (X *screen) Transparent() bool {
  if underX {
    return X.XWindow.Transparent()
  }
  return X.Console.Transparent()
}

func (X *screen) Transparence (on bool) {
  if underX {
    X.XWindow.Transparence (on)
  } else {
    X.Console.Transparence (on)
  }
}

func (X *screen) Write1 (b byte, l, c uint) {
  if underX {
    X.XWindow.Write1 (b, l, c)
  } else {
    X.Console.Write1 (b, l, c)
  }
}

func (X *screen) Write (s string, l, c uint) {
  if len(s) == 0 { return }
  if underX {
    X.XWindow.Write (str.Lat1 (s), l, c)
  } else {
    X.Console.Write (str.Lat1 (s), l, c)
  }
}

func (X *screen) WriteNat (n, l, c uint) {
  s:= "0"
  if n > 0 {
    s = ""
    for n > 0 {
      s = string(byte('0') + byte(n % 10)) + s
      n /= 10
    }
  }
  X.Write (s, l, c)
}

func (X *screen) WriteNatGr (n uint, x, y int) {
  s:= "0"
  if n > 0 {
    s = ""
    for n > 0 {
      s = string(byte('0') + byte(n % 10)) + s
      n /= 10
    }
  }
  X.WriteGr (s, x, y)
}

func (X *screen) Write1Gr (b byte, x, y int) {
  if underX {
    X.XWindow.Write1Gr (b, x, y)
  } else {
    X.Console.Write1Gr (b, x, y)
  }
}

func (X *screen) WriteGr (s string, x, y int) {
  if len(s) == 0 { return }
  if underX {
    X.XWindow.WriteGr (str.Lat1 (s), x, y)
  } else {
    X.Console.WriteGr (str.Lat1 (s), x, y)
  }
}

func (X *screen) Write1InvGr (b byte, x, y int) {
  if underX {
    X.XWindow.Write1InvGr (b, x, y)
  } else {
    X.Console.Write1InvGr (b, x, y)
  }
}

func (X *screen) WriteInvGr (s string, x, y int) {
  if len(s) == 0 { return }
  if underX {
    X.XWindow.WriteInvGr (str.Lat1 (s), x, y)
  } else {
    X.Console.WriteInvGr (str.Lat1 (s), x, y)
  }
}
