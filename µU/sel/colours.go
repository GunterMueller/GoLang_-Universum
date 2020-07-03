package sel

// (c) Christian Maurer   v. 191125 - license see µU.go

import (
  "µU/kbd"
  "µU/col"
  "µU/scr"
)
const
  H = 32
var (
  X, Y int
  pattern16 [M]col.Colour
  pattern [N]col.Colour
)

func init() {
  pattern16[ 0] = col.Black()
  pattern16[ 1] = col.Brown()
  pattern16[ 2] = col.Red()
  pattern16[ 3] = col.LightRed()
  pattern16[ 4] = col.Yellow()
  pattern16[ 5] = col.LightGreen()
  pattern16[ 6] = col.Green()
  pattern16[ 7] = col.Cyan()
  pattern16[ 8] = col.LightCyan()
  pattern16[ 9] = col.LightBlue()
  pattern16[10] = col.Blue()
  pattern16[11] = col.Magenta()
  pattern16[12] = col.LightMagenta()
  pattern16[13] = col.LightWhite()
  pattern16[14] = col.White()
  pattern16[15] = col.Gray()

  pattern[ 0] = col.DarkBrown()
  pattern[ 1] = col.Brown()
  pattern[ 2] = col.LightBrown()
  pattern[ 3] = col.DarkRed()
  pattern[ 4] = col.Red()
  pattern[ 5] = col.FlashRed()
  pattern[ 6] = col.LightRed()
  pattern[ 7] = col.DarkOrange()
  pattern[ 8] = col.Orange()
  pattern[ 9] = col.LightOrange()
  pattern[10] = col.Pink()
  pattern[11] = col.DarkYellow()
  pattern[12] = col.FlashYellow()
  pattern[13] = col.Yellow()
  pattern[14] = col.LightYellow()
  pattern[15] = col.LightGreen()
  pattern[16] = col.FlashGreen()
  pattern[17] = col.Green()
  pattern[18] = col.DarkGreen()
  pattern[19] = col.DarkCyan()
  pattern[20] = col.Cyan()
  pattern[21] = col.FlashCyan()
  pattern[22] = col.LightCyan()
  pattern[23] = col.DarkBlue()
  pattern[24] = col.Blue()
  pattern[25] = col.FlashBlue()
  pattern[26] = col.LightBlue()
  pattern[27] = col.DarkMagenta()
  pattern[28] = col.Magenta()
  pattern[29] = col.FlashMagenta()
  pattern[30] = col.LightMagenta()
  pattern[31] = col.Black()
  pattern[32] = col.DarkGray()
  pattern[33] = col.Gray()
  pattern[34] = col.LightGray()
  pattern[35] = col.White()
}

func write (FZ, B, x, y int) {
  f := scr.ColF()
  for i := 0; i < FZ; i++ {
    switch FZ { case 16:
      scr.ColourF (pattern16[i])
    default:
      scr.ColourF (pattern[i])
    }
    scr.RectangleFull (x + i * B, y, x + i * B + B - 1, y + H - 1)
  }
  scr.ColourF (f)
}

func define (FZ, B uint, C *col.Colour) {
  xi, yi := scr.MousePosGr()
  x, y := xi, yi
  x -= X
  x = x / int(B)
  if x < int(FZ) && Y <= y && y < Y + H {
    if FZ == M {
      *C = pattern16[x]
    } else {
      *C = pattern[x]
    }
    scr.ColourF (*C)
  } else {
    *C = scr.ScrColB()
  }
}

func colour16() col.Colour {
  return co (M)
}

func colour() col.Colour {
  return co (N)
}

func co (FZ uint) col.Colour {
  B := uint(M)
  if FZ == N { B = 32 }
  MausAn := scr.MousePointerOn()
  if ! MausAn {
    scr.MousePointer (true)
  }
  xm, ym := scr.MousePosGr()
  X, Y = xm, ym
  M := int(FZ * B) / 2
  if X >= int(scr.Ht()) - M { X = int(scr.Wd()) - M }
  if X >= M { X -= M } else { X = 0 }
  if Y >= H { Y -= H } else { Y = 0 }
  scr.SaveGr (X, Y, X + 2 * int(FZ * B), Y + H)
  write (int(FZ), int(B), X, Y)
  clicked := false
  C := scr.ScrColF()
  loop: for {
    scr.MousePointer (true)
    K, _ := kbd.Command()
    switch K {
    case kbd.Esc, kbd.Back, kbd.There, kbd.This:
      break loop
    case kbd.Here:
      define (FZ, B, &C)
      clicked = true
    case kbd.Hither:
      if clicked { break loop }
    }
  }
  scr.RestoreGr (X, Y, X + 2 * int(FZ * B), Y + H)
  if ! MausAn {
    scr.MousePointer (false)
  }
  return C
}
