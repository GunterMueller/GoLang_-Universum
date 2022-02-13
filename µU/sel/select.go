package sel

// (c) Christian Maurer   v. 220211 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/env"
  "µU/str"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
  . "µU/files"
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

var
  ptSuffix string

func hasSuffix (a Any) bool {
  p, ok := str.Sub (ptSuffix, a.(string))
  return ok && p == str.ProperLen (a.(string)) - uint(len (ptSuffix))
}

func aus (n, l, c uint, f, b col.Colour) {
//  N := files.NamePred (hasSuffix, n)
  N := NamePred (hasSuffix, n)
  if p, ok := str.Sub (ptSuffix, N); ok {
    N = str.Part (N, 0, p)
  }
  bx.Colours (f, b)
  bx.Write (N, l, c)
}

func names (mask, suffix string, n uint, l, c uint, f, b col.Colour) (string, string) {
  t, t1 := uint(len (mask)), uint(0)
  if t > 0 {
    t1 = 1 + t
  }
  scr.Save (l, c, t1 + n, 1)
  if t > 0 {
    mbx.Wd (t)
    mbx.ScrColours()
    mbx.Write (mask, l, c)
  }
  bx.Wd (n)
  bx.Colours (f, b)
  ptSuffix = "." + suffix
  errh.Hint ("falls Dateien vorhanden, auswählen F2-, dann Pfeil-/Eingabetaste, ggf. Esc")
  name := env.Arg(1)
  if name == "" {
    name = str.New (n) // Wörkeraunt um Fehler in box/imp.go
  }
  if p, ok := str.Pos (name, '.'); ok {
    name = str.Part (name, 0, p)
  }
  bx.Edit (&name, l, c + t1)
  str.OffSpc (&name)
  if p, ok := str.Pos (name, '.'); ok {
    name = str.Part (name, 0, p)
  }
  filename := name + ptSuffix
//  a := files.NumPred (hasSuffix)
  a := NumPred (hasSuffix)
  if a > 0 {
    switch C, _ := kbd.LastCommand(); C {
    case kbd.Esc:
      return "", "" // str.New (n), ""
    case kbd.Enter:
      // entered
    case kbd.Search:
      i := uint(0)
      select_ (aus, a, a, n, &i, l, c + t1, b, f)
      if i == a {
        return "", "" // str.New (n), ""
      } else {
//        filename = str.Lat1 (files.NamePred (hasSuffix, i))
        filename = str.Lat1 (NamePred (hasSuffix, i))
      }
    }
  }
  errh.DelHint()
  str.OffSpc (&filename)
  if p, ok := str.Pos (filename, '.'); ok {
    name = str.Part (filename, 0, p)
  }
  scr.Restore (l, c, t1 + n, 1)
  return name, filename
}
