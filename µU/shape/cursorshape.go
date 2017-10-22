package shape

// (c) Christian Maurer   v. 140204 - license see µU.go

func cursor (x, y, h uint, c, s Shape) (uint, uint) {
  if c == s { return 0, 0 }
  const s0 = 2
  var y0, y1 uint
  switch s { case Off:
    switch c { case Understroke:
      y0, y1 = h - s0, h - 1
    case Block:
      y0, y1 = 0, h - 1
    }
  case Understroke:
    switch c { case Off:
      y0, y1 = h - s0, h - 1
    case Block:
      y0, y1 = 0, h - s0 - 1
    }
  case Block:
    switch c { case Off:
      y0, y1 = 0, h - 1
    case Understroke:
      y0, y1 = 0, h - s0 - 1
    }
  }
  return y0, y1
}
