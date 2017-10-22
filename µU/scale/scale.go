package scale

// (c) Christian Maurer   v. 150425 - license see µU.go

import (
  "math"
  "µU/kbd"
  "µU/scr"
)
const (
  minInt = int(math.MinInt32)
  maxInt = int(math.MaxInt32)
)
var (
  nX, nY,              // pixel per screen
  x0, y0,              // left lower corner
  xMin, yMin,          // left limit
  xm, ym,              // center
  width, height,       //
  maxWidth, maxHeight, //
  maxMag,              //
  mX, mY float64       // transformation factors
  mm, nn [3]float64 = [3]float64 { math.Sqrt (math.Sqrt (2)), math.Sqrt (2), 2 }, // magnification factor
                      [3]float64 { 16, 4, 1 } // move part
)

func setRange (x, y, w float64) {
  nX, nY = float64(scr.Wd()), float64(scr.Ht())
  x0, y0 = x, y
  if w <= 0 {
    width = 1
  } else {
    width = w
  }
  height = width / scr.Proportion()
  xm, ym = x0 + width / 2, y0 + height / 2
  mX, mY = nX / width, nY / height
}

func scale (x, y float64) (int, int) {
  rx, ry:= int(math.Trunc (mX * (x - x0)) + 0.5), int(math.Trunc (mY * (y - y0)) + 0.5)
  ry = int(scr.Ht()) - ry
  if scr.UnderX() {
    if minInt <= rx && rx < maxInt && minInt <= ry && ry < maxInt {
      return rx, ry
    }
  } else if x0 <= x && x < x0 + width && y0 <= y && y < y0 + height {
    return rx, ry
  }
  return maxInt, maxInt
}

func rescale (x, y int) (float64, float64) {
  return x0 + float64(x) / mX, y0 + float64(int(scr.Ht()) - y) / mY
}

func lim (x, y, w, h, m float64) {
  xMin, yMin = x, y
  maxWidth, maxHeight, maxMag = w, h, m
}

var
  xt, yt float64 // to redefine globally !

func edit() {
  c, d:= kbd.LastCommand ()
  if d > 2 { d = 2 }
  switch c { case kbd.Back:
    if width < maxWidth {
      w:= width * mm [d]
      if w > maxWidth {
        w = maxWidth
      }
      dw:= (w - width) / 2
      x0, y0 = x0 - dw, y0 - dw / scr.Proportion()
      width = w
    } else { // überschritten
      x0, y0 = xMin, yMin
      width = maxWidth
    }
    height = width / scr.Proportion()
    mX, mY = nX / width, nY / height
  case kbd.Enter:
    if width > maxWidth / maxMag {
      w:= width / mm [d]
      if w < maxWidth / maxMag {
        w = maxWidth / maxMag
      }
      dw:= (w - width) / 2
      x0, y0 = x0 - dw, y0 - dw / scr.Proportion()
      width = w
      height = width / scr.Proportion()
      mX, mY = nX / width, nY / height
    }
  case kbd.Left:
    if x0 >= xMin {
      x0 = x0 - width / nn [d]
    }
  case kbd.Right:
    if x0 + width <= xMin + maxWidth {
      x0 = x0 + width / nn [d]
    }
// TODO: Rollrad von Maus einbauen - die sendet Up/Down
  case kbd.Up:
    if y0 + height < yMin + maxHeight {
      y0 = y0 + width / nn [d]
    }
  case kbd.Down:
    if y0 >= yMin {
      y0 = y0 - width / nn [d]
    }
  case kbd.This:
    x, y:= scr.MousePosGr()
    y = int(scr.Ht()) - y
    x0 = x0 + float64 (x) / mX - width / 2
    y0 = y0 + float64 (y) / mY - height / 2
  case kbd.There:
    x, y:= scr.MousePosGr()
    xt, yt = float64 (x), float64 (int(scr.Ht()) - y)
  case kbd.Push, kbd.Thither:
    x, y:= scr.MousePosGr()
    x0 = x0 - (float64 (x) - xt) / mX
    y0 = y0 - (float64 (int(scr.Ht()) - y) - yt) / mY
    xt, yt = float64 (x), float64 (int(scr.Ht()) - y)
  }
  if x0 < xMin {
    x0 = xMin
  }
  if x0 + width > xMin + maxWidth {
    x0 = xMin + maxWidth - width
  }
  if y0 < yMin {
    y0 = yMin
  }
  if y0 + height >= yMin + maxHeight {
    y0 = yMin + maxHeight - height
  }
}

func init_() {
  setRange (0, 0, 1) // must not be called in init, because it uses initial data of the screen !
  lim (0, 0, width, height, 1)
}
