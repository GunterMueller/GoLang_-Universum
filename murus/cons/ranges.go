package cons

// (c) Christian Maurer   v. 140527 - license see murus.go

import (
  "strconv"
  "murus/col"
)

func (X *console) rectangOk (x, y, x1, y1 *int) bool {
  if ! visible { return false }
  intord (x, y, x1, y1)
  if *x >= int(X.wd) || *y >= int(X.ht) {
    return false
  }
  if *x1 >= int(X.wd) { *x1 = int(X.wd) - 1 }
  if *y1 >= int(X.ht) { *y1 = int(X.ht) - 1 }
  return true
}

func (X *console) urectangOk (x, y, x1, y1 *uint) bool {
  if ! visible { return false }
  if *x > *x1 { *x, *x1 = *x1, *x }
  if *y > *y1 { *y, *y1 = *y1, *y }
  if *x >= X.wd || *y >= X.ht {
    return false
  }
  if *x1 >= X.wd { *x1 = X.wd - 1 }
  if *y1 >= X.ht { *y1 = X.ht - 1 }
  return true
}

func (X *console) Cls() {
  if ! visible { return }
  l := int(colourdepth) * int(X.wd)
  a := 0
  c := col.Cc (X.scrB)
  for j := 0; j < int(X.ht); j++ {
    for i := 0; i < int(X.wd); i++ {
      copy (emptyBackground[a:a+int(colourdepth)], c)
      a += int(colourdepth)
    }
  }
  a = (X.y * int(width) + X.x) * int(colourdepth)
  for j := 0; j < int(X.ht); j++ {
    copy (fbmem[a:a+l], emptyBackground)
    copy (fbcop[a:a+l], emptyBackground)
    a += int(width) * int(colourdepth)
  }
}

func u (n uint) string { return strconv.Itoa(int(n)) }
func i (n  int) string { return strconv.Itoa(    n ) }

func (X *console) Clr (l, c, w, h uint) {
  x, y := int(c * X.wd1), int(l * X.ht1)
  X.ClrGr (x, y, x + int(w * X.wd1), y + int(h * X.ht1))
}

func (X *console) ClrGr (x, y, x1, y1 int) {
  if ! X.rectangOk (&x, &y, &x1, &y1) { return }
  if ! visible { return }
  x += X.x; x1 += X.x // y's diff !
  da := uint(x1 - x) * colourdepth
  a := uint(0)
/*
  c := col.Cc (X.ScrColB())
  for j := uint(0); j < da; j++ {
    copy (emptyBackground[a:a+colourdepth], c)
    a += colourdepth
  }
*/
  w := width * colourdepth
  a = uint(X.y + y) * w + uint(x) * colourdepth
  for z := 0; z <= y1 - y; z++ {
    copy (fbmem[a:a+da], emptyBackground[:da])
    copy (fbcop[a:a+da], emptyBackground[:da])
    a += w
  }
}

func (X *console) Buf (on bool) {
  if X.buff == on { return }
  X.buff = on
/*
  a := 0
  c := col.Cc (X.ScrColB())
  for x := 0; x < int(X.wd); x++ {
    copy (emptyBackground[a:a+int(colourdepth)], c)
    a += int(colourdepth)
  }
*/
  da := int(X.wd) * int(colourdepth)
  w := int(width) * int(colourdepth)
  a := (int(width) * X.y + X.x) * int(colourdepth)
  for y := 0; y < int(X.ht); y++ {
    if on {
      copy (fbcop[a:a+da], emptyBackground[:da])
    } else {
      copy (fbmem[a:a+da], fbcop[a:a+da])
    }
    a += w
  }
}

func (x *console) Buffered () bool {
  return x.buff
}

func Buf1 (on bool) {
  for _, s := range consList {
    s.buff = on
  }
/*
  c := col.Cc (col.Black)
  a := 0
  for i := 0; i < int(width); i++ {
    for j := 0; j < int(height); j++ {
      copy (emptyBackground[a:a+int(colourdepth)], c)
      a += int(colourdepth)
    }
  }
*/
  if on {
    copy (fbcop, emptyBackground)
  } else {
    copy (fbmem, fbcop)
  }
}

func (X *console) Save (l, c, w, h uint) {
  x, y := int(X.wd1 * c), int(X.ht1 * l)
  X.SaveGr (x, y, x + int(X.wd1 * w), y + int(X.ht1 * h))
}

func (X *console) SaveGr (x, y, x1, y1 int) {
  if ! X.rectangOk (&x, &y, &x1, &y1) { return }
  w, h := x1 - x + 1, y1 - y + 1
  x0, y0 := X.x + x, X.y + y
  a, da := x * int(colourdepth), w * int(colourdepth)
  if X.mouseOn { X.MousePointer (false) }
  for i := 0; i < h; i++ {
    b := (int(width) * (y0 + i) + x0) * int(colourdepth)
    copy (X.shadow[i][a:a+da], fbmem[b:b+da])
  }
  if X.mouseOn { X.MousePointer (true) }
}

func (X *console) Save1 () {
  X.SaveGr (0, 0, int(X.wd) - 1, int(X.ht) - 1)
}

func (X *console) Restore (l, c, w, h uint) {
  x, y := int(X.wd1 * c), int(X.ht1 * l)
  X.RestoreGr (x, y, x + int(X.wd1 * w), y + int(X.ht1 * h))
}

func (X *console) RestoreGr (x, y, x1, y1 int) {
  if ! X.rectangOk (&x, &y, &x1, &y1) { return }
  w, h := x1 - x + 1, y1 - y + 1
  x0, y0 := X.x + x, X.y + y
  a, da := x * int(colourdepth), w * int(colourdepth)
  for i := 0; i < h; i++ {
    b := (int(width) * (y0 + i) + x0) * int(colourdepth)
    copy (fbmem[b:b+da], X.shadow[i][a:a+da])
    copy (fbcop[b:b+da], X.shadow[i][a:a+da])
  }
}

func (X *console) Restore1() {
  X.RestoreGr (0, 0, int(X.wd) - 1, int(X.ht) - 1)
}
