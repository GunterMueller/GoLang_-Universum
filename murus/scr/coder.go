package scr

// (c) murus.org  v. 140527 - license see murus.go

func (X *screen) Codelen (w, h uint) uint {
  if underX {
    return X.Window.Codelen (w, h)
  }
  return X.Console.Codelen (w, h)
}

func (X *screen) Encode (x, y, w, h uint) []byte {
  if underX {
    return X.Window.Encode (x, y, w, h)
  }
  return X.Console.Encode (x, y, w, h)
}

func (X *screen) Decode (bs []byte) {
  if underX {
    X.Window.Decode (bs)
  } else {
    X.Console.Decode (bs)
  }
}
