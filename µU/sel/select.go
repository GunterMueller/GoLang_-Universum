package sel

// (c) Christian Maurer   v. 220815 - license see µU.go

import (
  "µU/ker"
  "µU/kbd"
  "µU/str"
  "µU/col"
  "µU/font"
  "µU/scr"
  "µU/box"
)
var
  bx, mbx = box.New(), box.New()

func select_ (write WritingCol, n, h, w uint, i *uint, l, c uint, f, b col.Colour) {
  if n == 0 { ker.Oops() }
  if n == 1 { *i = 0; return }
  if h == 0 { ker.Oops() }
  if h > n { h = n }
  if w == 0 { w = scr.NColumns() }
  if w > scr.NColumns() { w = scr.NColumns() }
  if c + w > scr.NColumns() { c = scr.NColumns() - w }
// so, that last line remains free
  if l + h >= scr.NLines() {
    h = scr.NLines() - l - 1
  }
  if *i >= n { *i = n - 1 }
  MouseOn := scr.MousePointerOn()
  var x, y int
  if MouseOn {
    scr.MousePointer (false)
    x, y = scr.MousePosGr()
  }
  scr.WarpMouse (l + *i, c)
  scr.Save (l, c, w, h)
  i0, n0 := uint(0), uint(0)
  if *i == 0 { n0 = 1 } // else { n0 = 0 }
  neu := true
  loop:
  for {
    if *i < i0 {
      i0 = *i
      neu = true
    } else if *i > i0 + h - 1 {
      i0 = *i - (h - 1)
      neu = true
    } else {
      neu = *i != n0
    }
    if neu {
      neu = false
      var cF, cB col.Colour
      for j := uint(0); j < h; j++ {
        if i0 + j == *i {
          cF, cB = f, b
        } else {
          cF, cB = b, f
        }
        write (i0 + j, l + j, c, cF, cB)
      }
    }
    n0 = *i
    C, d := kbd.Command()
    switch C {
    case kbd.Esc, kbd.Move:
      *i = n
      break loop
    case kbd.Enter, kbd.Here:
      break loop
    case kbd.Left, kbd.Up, kbd.ScrollUp:
      if d == 0 {
        if *i > 0 {
          *i--
        }
      } else {
        if *i >= 10 {
          *i -= 10
        }
      }
    case kbd.Right, kbd.Down, kbd.ScrollDown:
      if d == 0 {
        if *i + 1 < n {
          *i++
        }
      } else {
        if *i + 10 < n {
          *i += 10
        }
      }
    case kbd.Pos1:
      *i = 0
    case kbd.End:
      *i = n - 1
    case kbd.Go:
      _, yM := scr.MousePosGr()
      if uint(yM) <= l * scr.Ht1() + scr.Ht1() / 2 {
        if *i > 0 {
          *i --
        }
      } else if uint(yM) >= (l + h) * scr.Ht1() {
        if *i < n - 1 {
          *i ++
        }
      } else {
        *i = i0 + uint(yM) / scr.Ht1() - l
      }
/*/
    case kbd.Help:
      errh.Hint (errh.ToSelect)
      kbd.Wait (true)
      errh.DelHint()
/*/
    }
  }
  scr.Restore (l, c, w, h)
  if MouseOn {
    scr.MousePointer (true)
    scr.WarpMouseGr (x, y)
  }
}

func select1 (s []string, h, w uint, n *uint, l, c uint, f, b col.Colour) {
  ls := uint(len(s))
//  if ls + 1 < h { h = ls + 1 }
  bx.Wd (w)
  select_(func (k, l, c uint, f, b col.Colour) {
            if k < ls { bx.Colours (f, b); bx.Write (s[k], l, c) }
          }, h, h, w, n, l, c, f, b)
}

func colour (l, c, w uint) (col.Colour, bool) {
  return colours (l, c, w, col.Black(), col.DarkRed(), col.Red(), col.FlashRed(),
                           col.LightRed(), col.FlashOrange(), col.DarkYellow(), col.Yellow(),
                           col.FlashGreen(), col.Green(), col.DarkGreen(), col.DarkCyan(),
                           col.Cyan(), col.FlashCyan(), col.LightBlue(), col.FlashBlue(),
                           col.Blue(), col.Magenta(), col.FlashMagenta(), col.LightMagenta(),
                           col.White(), col.LightGray(), col.Gray(), col.DarkGray(),
                           col.DarkBrown(), col.Brown(), col.LightBrown(), col.LightWhite())
}

func colours (l, c, w uint, cols ...col.Colour) (col.Colour, bool) {
  scr.Save (l, c, w, 1)
  s := str.New (w)
  n, i := uint(len(cols)), uint(0)
  loop:
  for {
    scr.ColourB (cols[i])
    scr.Write (s, l, c)
    switch cmd, _ := kbd.Command(); cmd {
    case kbd.Esc:
      scr.Restore (l, c, w, 1)
      return col.Black(), false
    case kbd.Enter:
      break loop
    case kbd.Up:
      if i > 0 {
        i--
      }
    case kbd.Down:
      if i + 1 < n {
        i++
      }
    case kbd.Pos1:
      i = 0
    case kbd.End:
      i = n - 1
    }
  }
  scr.Restore (l, c, w, 1)
  return cols[i], true
}

func fontsize (f, b col.Colour) font.Size {
  n := uint(0)
  scr.MousePointer (true)
  z, s := scr.MousePos()
  Select1 (font.Name, uint(font.NSizes), font.M, &n, z, s, f, b)
  if n < uint(font.NSizes) {
    return font.Size (n)
  }
  return font.Normal
}
