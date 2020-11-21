package cons

// (c) Christian Maurer   v. 201102 - license see µU.go

import (
  "math"
  . "µU/linewd"
  "µU/ker"
)
type
  pointFunc func (int, int)

func (X *console) SetLinewidth (w Linewidth) {
  X.lineWd = w
}

func (X *console) ActLinewidth() Linewidth {
  return X.lineWd
}

func (X *console) iok (x, y int) bool {
  if ! visible { return false }
  if x < 0 || y < 0 { return false }
  return x < int(X.wd) && y < int(X.ht)
//  return x < X.x + int(X.wd) && y < X.y + int(X.ht)
}

func (X *console) iok4 (x, y, x1, y1 int) bool {
  if ! visible { return false }
  if x < 0 || y < 0 || x1 < 0 || y1 < 0 { return false }
  return true
  return x < int(X.wd) && y < int(X.ht) && x1 < int(X.wd) && y1 < int(X.ht) // shit
}

func intord (x, y, x1, y1 *int) {
  if *x > *x1 { *x, *x1 = *x1, *x }
  if *y > *y1 { *y, *y1 = *y1, *y }
}

func (X *console) Point (x, y int) {
  if ! visible || ! X.iok (x, y) { return }
  x += X.x; y += X.y
//  ux, uy := uint(x), uint(y)
  a := (int(width) * y + x) * int(colourdepth)
  copy (fbcop[a:a+int(colourdepth)], X.codeF)
  if ! X.buff {
    copy (fbmem[a:a+int(colourdepth)], X.codeF)
  }
/*
  if X.lineWd > Thin && ux + 1 < X.wd && uy + 1 < X.ht {
    if ux + 1 < X.ht {
      a += int(colourdepth)
      copy (fbcop[a:a+int(colourdepth)], X.codeF)
      if ! X.buff {
        copy (fbmem[a:a+int(colourdepth)], X.codeF)
      }
    }
    if uy + 1 < X.wd {
      a += int(width - 1) * int(colourdepth)
      copy (fbcop[a:a+int(colourdepth)], X.codeF)
      if ! X.buff {
        copy (fbmem[a:a+int(colourdepth)], X.codeF)
      }
    }
    if X.lineWd == Thick {
      a += int(colourdepth)
      copy (fbcop[a:a+int(colourdepth)], X.codeF)
      if ! X.buff {
        copy (fbmem[a:a+int(colourdepth)], X.codeF)
      }
    } else { // Thicker
      if ux > 0 && uy > 0 {
        a -= int(width * 2 * colourdepth)
        copy (fbcop[a:a+int(colourdepth)], X.codeF)
        if ! X.buff {
          copy (fbmem[a:a+int(colourdepth)], X.codeF)
        }
        a += int(width - 1) * int(colourdepth)
        copy (fbcop[a:a+int(colourdepth)], X.codeF)
        if ! X.buff {
          copy (fbmem[a:a+int(colourdepth)], X.codeF)
        }
      }
    }
  }
  if X.lineWd > Thin && ux + 1 < X.wd && uy + 1 < X.ht { // still buggy TODO
    a += int(colourdepth)
    copy (fbcop[a:a+int(colourdepth)], X.codeF)
    a += int(width - 1) * int(colourdepth)
    copy (fbcop[a:a+int(colourdepth)], X.codeF)
    if X.lineWd == Thick {
      a += int(colourdepth)
      copy (fbcop[a:a+int(colourdepth)], X.codeF)
    } else { // Thicker
      if ux > 0 && uy > 0 {
        a -= int(width * 2 * colourdepth)
        copy (fbcop[a:a+int(colourdepth)], X.codeF)
        a += int(width - 1) * int(colourdepth)
        copy (fbcop[a:a+int(colourdepth)], X.codeF)
      }
    }
  }
*/
}

func (X *console) PointInv (x, y int) {
  if ! X.iok (x, y) { return }
  c := X.Colour (uint(x), uint(y))
  c.Invert()
  X.ColourF (c)
  X.Point (x, y)
  X.ColourF (X.cF)
}

func (X *console) OnPoint (x, y, a, b int, d uint) bool {
  dx, dy := x - a, y - b
  return dx * dx + dy * dy <= int(d * d)
}

// Returns true, iff m is up to tolerance t between i and k.
func between (i, k, m, t int) bool {
  return i <= m + t && m <= k + t || k <= m + t && m <= i + t
}

func ok2 (xs, ys []int) bool {
  if ! visible { return false }
  n := len (xs)
  return n != 0 && n == len (ys)
}

func ok4 (xs, ys, xs1, ys1 []int) bool {
  if ! visible { return false }
  n := len (xs)
  return n != 0 && n == len (ys) && n == len (xs1) && len (xs1) == len (ys1)
}

func (X *console) Points (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  for i := 0; i < len (xs); i++ {
    X.Point (xs[i], ys[i])
  }
}

func (X *console) PointsInv (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  for i := 0; i < len (xs); i++ {
    X.PointInv (xs[i], ys[i])
  }
}

// Pre: x <= x1 < Wd, y < Ht.
func (X *console) horizontal (x, y, x1 int, f pointFunc) {
  if x == x1 { f (x, y); return }
  if x > x1 { x, x1 = x1, x }
//  if x >= X.wd { return }
//  if x1 >= int(X.wd) { x1 = int(X.wd) - 1 }
  x0 := x
  for x := x0; x <= x1; x++ {
    f (x, y)
  }
/*
  if X.lineWd > Thin && y + 1 <= int(X.ht) {
    for x := x0; x <= x1; x++ {
      f (x, y + 1)
    }
  }
  if X.lineWd > Thick && y > 0 {
    for x := x0; x <= x1; x++ {
      f (x, y - 1)
    }
  }
*/
}

// Pre: x < Wd, y <= y1 < Ht.
func (X *console) vertical (x, y, y1 int, f pointFunc) {
  if y > y1 { y, y1 = y1, y }
//  if y1 >= int(X.ht) { y1 = int(X.ht) - 1 }
  y0 := y
  for y := y0; y <= y1; y++ {
    f (x, y)
  }
/*
  if X.lineWd > Thin && x + 1 < int(X.wd) {
    for y := y0; y <= y1; y++ {
      f (x + 1, y)
    }
  }
  if X.lineWd > Thick && x > 0 {
    for y := y0; y <= y1; y++ {
      f (x - 1, y)
    }
  }
*/
}

// Pre: 0 <= x <= x1 < NColumns, 0 <= y != y1 < NLines.
func (X *console) bresenham (x, y, x1, y1 int, f pointFunc) {
  dx := x1 - x
  Fehler, dy := 0, 0
  if y <= y1 { // Steigung positiv
    dy = y1 - y
    if dy <= dx { // Steigung <= 45°
      for {
        f (x, y)
        if x == x1 { break }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y++
          Fehler -= 2 * dx
        }
      }
    } else { // Steigung > 45°
      for {
        f (x, y)
        if y == y1 { break }
        y++
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
    }
  } else { // Steigung negativ
    dy = y - y1
    if dy <= dx { // Steigung >= -45°
      for {
        f (x, y)
        if x == x1 { break }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y--
          Fehler -= 2 * dx
        }
      }
    } else { // Steigung < -45°
      for {
        f (x, y)
        if y == y1 { break }
        y--
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
    }
  }
}

// Pre: 0 <= x <= x1 < xx, y != y1, 0 <= y, y1 < yy.
func (X *console) bresenhamInf (xx, yy, x, y, x1, y1 int, f pointFunc) {
  x0, y0 := x, y
  dx := x1 - x
  Fehler, dy := 0, 0
  if y <= y1 { // Steigung positiv
    dy = y1 - y
    if dy <= dx { // Steigung <= 45°
      for {
        f (x, y)
        if x == xx - 1 || y == yy - 1 { break }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y++
          Fehler -= 2 * dx
        }
      }
      x, y = x0, y0
      Fehler = 0
      for {
        Fehler += 2 * dy
        if Fehler > dx {
          y--
          Fehler -= 2 * dx
        }
        f (x, y)
        if x == 0 || y == 0 { break }
        x--
      }
    } else { // Steigung > 45°
      for {
        f (x, y)
        if y == yy - 1 || x == xx - 1 { break }
        y++
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
      x, y = x0, y0
      Fehler = 0
      for {
        Fehler += 2 * dx
        if Fehler > dy {
          x--
          Fehler -= 2 * dy
        }
        f (x, y)
        if x == 0 || y == 0 { break }
        y--
      }
    }
  } else {
    dy = y - y1 // Steigung negativ
    if dy <= dx { // Steigung >= -45°
      for {
        f (x, y)
        if x == xx - 1 || y == 0 { break }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y--
          Fehler -= 2 * dx
        }
      }
      x, y = x0, y0
      Fehler = 0
      for {
        f (x, y)
        if x == 0 || y == yy - 1 { break }
        x--
        Fehler += 2 * dy
        if Fehler > dx {
          y++
          Fehler -= 2 * dx
        }
      }
    } else { // Steigung < -45°
      for {
        f (x, y)
        if x == xx - 1 || y == 0 { break }
        y--
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
      x, y = x0, y0
      Fehler = 0
      for {
        f (x, y)
        if x == 0 || y == yy - 1 { break }
        y++
        Fehler += 2 * dx
        if Fehler > dy {
          x--
          Fehler -= 2 * dy
        }
      }
    }
  }
}

func nat (x, y int) bool {
  return x >= 0 && y >= 0
}

func (X *console) line (x, y, x1, y1 int, f pointFunc) {
  if x1 < x { x, x1 = x1, x; y, y1 = y1, y }
  if ! X.iok4 (x, y, x1, y1) {
    return
  }
  if y == y1 {
    X.horizontal (x, y, x1, f)
    return
  }
  if x == x1 {
    X.vertical (x, y, y1, f)
    return
  }
  X.bresenham (x, y, x1, y1, f)
}

func (X *console) Line (x, y, x1, y1 int) {
  X.line (x, y, x1, y1, X.Point)
}

func (X *console) LineInv (x, y, x1, y1 int) {
  X.line (x, y, x1, y1, X.PointInv)
}

func (X *console) OnLine (x, y, x1, y1, a, b int, t uint) bool {
  if x1 < x { x, x1 = x1, x; y, y1 = y1, y }
  if ! (between (x, x1, a, int(t)) && between (y, y1, b, int(t))) {
    return false
  }
  if x == x1 {
    return between (x, x, a, int(t)) && between (y, y1, b, int(t))
  }
  if y == y1 {
    return between (y, y, b, int(t)) && between (x, x1, a, int(t))
  }
  X.xx_, X.yy_, X.tt_, X.incident = a, b, int(t * t), false
  X.bresenham (x, y, x1, y1, X.onPoint)
  return X.incident
}

func (X *console) lines (xs, ys, xs1, ys1 []int, f pointFunc) {
  if ! ok4 (xs, ys, xs1, ys1) { return }
  for i := 0; i < len (xs); i++ {
    if X.iok (xs[i], ys[i]) && X.iok (xs1[i], ys1[i]) {
      X.line (xs[i], ys[i], xs1[i], ys1[i], f)
    }
  }
}

func (X *console) Lines (xs, ys, xs1, ys1 []int) {
  X.lines (xs, ys, xs1, ys1, X.Point)
}

func (X *console) LinesInv (xs, ys, xs1, ys1 []int) {
  X.lines (xs, ys, xs1, ys1, X.PointInv)
}

func (X *console) OnLines (xs, ys, xs1, ys1 []int, a, b int, t uint) bool {
  if ! ok4 (xs, ys, xs1, ys1) { return false }
  if len (xs) == 1 {
    return between (xs[0], xs[0], a, int(t)) && between (ys[0], ys[0], b, int(t))
  }
  for i := 0; i < len (xs); i++ {
    if X.OnLine (xs[i], ys[i], xs1[i], ys1[i], a, b, t) {
      return true
    }
  }
  return false
}

func (X *console) segs (xs, ys []int, f pointFunc) {
  if ! ok2 (xs, ys) { return }
  n := len (xs)
  for i := 0; i < n; i++ {
    if ! X.iok (xs[i], ys[i]) {
      return
    }
  }
  if n == 0 {
    f (xs[0], ys[0])
  } else {
    for i := 1; i < len (xs); i++ {
      X.line (xs[i-1], ys[i-1], xs[i], ys[i], f)
    }
  }
}

func (X *console) Segments (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  X.segs (xs, ys, X.Point)
}

func (X *console) SegmentsInv (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  X.segs (xs, ys, X.PointInv)
  if len (xs) > 1 {
    for i := 1; i < len (xs); i++ {
      X.PointInv (xs[i], ys[i])
    }
  }
}

func (X *console) OnSegments (xs, ys []int, a, b int, t uint) bool {
  if ! ok2 (xs, ys) { return false }
  if len (xs) == 1 {
    return xs[0] == a && ys[0] == b // TODO, weil das noch Blödsinn ist
  }
  for i := 1; i < len (xs); i++ {
    if X.OnLine (xs[i-1], ys[i-1], xs[i], ys[i], a, b, t) {
      return true
    }
  }
  return false
}

func (X *console) onPoint (x, y int) {
  X.incident = X.incident || (x - X.xx_) * (x - X.xx_) + (y - X.yy_) * (y - X.yy_) <= X.tt_
}

func (X *console) infLine (x, y, x1, y1 int, f pointFunc) {
  if x == x1 && y == y1 { return }
  if x1 < x { x, x1 = x1, x; y, y1 = y1, y }
  if ! visible { return }
  if y == y1 {
    X.horizontal (0, y, int(width) - 1, f)
    return
  }
  if x == x1 {
    X.vertical (x, 0, int(height), f)
    return
  }
  X.bresenhamInf (int(width), int(height), x, y, x1, y1, f)
}

func (X *console) InfLine (x, y, x1, y1 int) {
  X.infLine (x, y, x1, y1, X.Point)
}

func (X *console) InfLineInv (x, y, x1, y1 int) {
  X.infLine (x, y, x1, y1, X.PointInv)
}

func (X *console) OnInfLine (x, y, x1, y1, a, b int, t uint) bool {
  if x1 < x { x, x1 = x1, x; y, y1 = y1, y }
  X.xx_, X.yy_, X.tt_, X.incident = a, b, int(t * t), false
  X.bresenhamInf (int(width), int(height), x, y, x1, y1, X.onPoint)
  return X.incident
}

func (X *console) Triangle (x, y, x1, y1, x2, y2 int) {
  X.Line (x, y, x1, y1)
  X.Line (x1, y1, x2, y2)
  X.Line (x2, y2, x, y)
}

func (X *console) TriangleInv (x, y, x1, y1, x2, y2 int) {
  X.LineInv (x, y, x1, y1)
  X.LineInv (x1, y1, x2, y2)
  X.LineInv (x2, y2, x, y)
}

func (X *console) TriangleFull (x, y, x1, y1, x2, y2 int) {
  X.PolygonFull ([]int{x, x1, x2}, []int{y, y1, y2})
}

func (X *console) TriangleFullInv (x, y, x1, y1, x2, y2 int) {
  X.PolygonFullInv ([]int{x, x1, x2}, []int{y, y1, y2})
}

func (X *console) rectang (x, y, x1, y1 int, f pointFunc) {
  if ! X.rectangOk (&x, &y, &x1, &y1) { return }
  if x == x1 {
    if y == y1 {
      f (x, y)
    } else {
      X.vertical (int(x), int(y), int(y1), f)
    }
    return
  }
  X.horizontal (x, y, x1, f)
  if y == y1 {
    return
  }
  X.horizontal (x, y1, x1, f)
  X.vertical (x, y, y1, f)
  X.vertical (x1, y, y1, f)
}

func (X *console) Rectangle (x, y, x1, y1 int) {
  X.rectang (x, y, x1, y1, X.Point)
}

func (X *console) RectangleInv (x, y, x1, y1 int) {
  X.rectang (x, y, x1, y1, X.PointInv)
  X.PointInv (x, y)
  X.PointInv (x1, y)
  X.PointInv (x, y1)
  X.PointInv (x1, y1)
}

func (X *console) RectangleFull (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  if x1 >= int(X.wd) { x1 = int(X.wd) - 1 }
  if y1 >= int(X.ht) { y1 = int(X.ht) - 1 }
  for y <= y1 {
    X.horizontal (x, y, x1, X.Point)
    y++
  }
}

func (X *console) RectangleFullInv (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  if x1 >= int(X.wd) { x1 = int(X.wd) - 1 }
  if y1 >= int(X.ht) { y1 = int(X.ht) - 1 }
  for y <= y1 {
    X.horizontal (x, y, x1, X.PointInv)
    y++
  }
}

func (X *console) OnRectangle (x, y, x1, y1, a, b int, t uint) bool {
  if ! (between (x, x1, a, int(t)) && between (y, y1, b, int(t))) {
    return false
  }
  return between (x, x, a, int(t)) || between (x1, x1, a, int(t)) ||
         between (y, y, b, int(t)) || between (y1, y1, b, int(t))
}

func (X *console) InRectangle (x, y, x1, y1, a, b int, t uint) bool {
  return between (x, x1, a, int(t)) && between (y, y1, b, int(t))
}

func (X *console) Polygon (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  X.segs (xs, ys, X.Point)
  n := len (xs)
  if n > 1 {
    X.line (xs[n-1], ys[n-1], xs[0], ys[0], X.Point)
  }
}

func (X *console) PolygonInv (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  X.segs (xs, ys, X.PointInv)
  n := len (xs)
  if n > 1 {
    X.line (xs[n-1], ys[n-1], xs[0], ys[0], X.PointInv)
    X.PointInv (xs[0], ys[0])
    X.PointInv (xs[n-1], ys[n-1])
  }
}

func (X *console) interior (x, y int, xs, ys []int) bool {
  return false // TODO winding number algorithm
}

func (X *console) mark (x, y int) {
  X.polygon[x][y] = true
  X.Point (x, y)
}

func (X *console) markInv (x, y int) {
  X.polygon[x][y] = true
  X.PointInv (x, y)
}

func (X *console) demark (x, y int) {
  X.polygon[x][y] = false
}

func (X *console) dedone() {
  for x := uint(0); x < X.wd; x++ {
    for y := uint(0); y < X.ht; y++ {
      X.done[x][y] = false
    }
  }
}

func (X *console) st (x, y int, f pointFunc) {
  if X.polygon[x][y] {
    return
  }
  if ! X.done[x][y] {
    X.done[x][y] = true
    f (x, y)
    if y > 0 { X.st (x, y - 1, f) }
    if x > 0 { X.st (x - 1, y, f) }
    if y + 1 < int(X.ht) { X.st (x, y + 1, f) }
    if x + 1 < int(X.wd) { X.st (x + 1, y, f) }
  }
}

func (X *console) setInv (x, y int) {
  X.st (x, y, X.PointInv)
}

func (X *console) set (x, y int) {
  X.st (x, y, X.Point)
}

func (X *console) polygonFull (xs, ys []int, m, s pointFunc) {
  if ! ok2 (xs, ys) { return }
  n := len (xs)
  if n < 2 { return }
  X.segs (xs, ys, m)
  xx, yy := 0, 0
  xMin, yMin := int(X.wd), int(X.ht)
  xMax, yMax := 0, 0
  for i := 0; i < int(n); i++ {
    xx += xs[i]; yy += ys[i]
    if xs[i] < xMin { xMin = xs[i] }
    if ys[i] < yMin { yMin = ys[i] }
    if xs[i] > xMax { xMax = xs[i] }
    if ys[i] > yMax { yMax = ys[i] }
  }
  s (xx / n, yy / n)
  X.segs (xs, ys, X.demark)
  X.dedone()
}

func (X *console) PolygonFull (xs, ys []int) {
  X.polygonFull (xs, ys, X.mark, X.set)
}

func (X *console) PolygonFullInv (xs, ys []int) {
  X.polygonFull (xs, ys, X.markInv, X.setInv)
}

func (X *console) OnPolygon (xs, ys []int, a, b int, t uint) bool {
  n := len (xs)
  if n == 0 { return false }
  if ! ok2 (xs, ys) { return false }
  if n == 1 { return xs[0] == a && ys[0] == b }
  for i := 1; i < int(n); i++ {
    if X.OnLine (xs[i-1], ys[i-1], xs[i], ys[i], a, b, t) {
      return true
    }
  }
  return X.OnLine (xs[n-1], ys[n-1], xs[0], ys[0], a, b, t)
}

func (X *console) circ (x, y int, r uint, filled bool, f pointFunc) {
// Algorithmus von Bresenham (Fellner: Computer Grafik, 5.5)
  if ! visible { return }
  if x >= int(X.wd) || y >= int(X.ht) || r >= X.wd {
    return
  }
  if r == 0 {
    f (x, y)
    return
  }
  x1, y1 := 0, int(r)
  Fehler := 3
  Fehler -= 2 * int(r)
/*
  if filled {
    X.horizontal (x - r, y, x + r, b)
    X.Point (x, y - r)
    X.Point (x, y + r)
  } else {
    f (x - r, y    )
    f (x + r, y    )
    f (x    , y - r)
    f (x    , y + r)
  }
  x1++
  if Fehler >= 0 {
    y1--
    Fehler -= 4 * y1
  }
  Fehler += 6
*/
  y0 := y1 + 1
  for x1 <= y1 {
    if filled {
      X.horizontal (x - y1, y - x1, x + y1, f)
      if x1 > 0 {
        X.horizontal (x - y1, y + x1, x + y1, f)
      }
      if y1 < y0 { // not yet correct, but a bit better than the above code
        y0 = y1
        X.horizontal (x - x1, y - y1, x + x1, f)
        X.horizontal (x - x1, y + y1, x + x1, f)
      }
    } else {
      f (x - y1, y - x1)
      f (x + y1, y - x1)
      f (x - y1, y + x1)
      f (x + y1, y + x1)
      f (x - x1, y - y1)
      f (x + x1, y - y1)
      f (x - x1, y + y1)
      f (x + x1, y + y1)
    }
    x1++
    if Fehler >= 0 {
      y1--
      Fehler -= 4 * y1
    }
    Fehler += 4 * x1 + 2
  }
}

func (X *console) Circle (x, y int, r uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.circ (x, y, r, false, X.Point)
  }
}

func (X *console) CircleInv (x, y int, r uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.circ (x, y, r, false, X.PointInv)
  }
}

func (X *console) CircleFull (x, y int, r uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.circ (x, y, r, true, X.Point)
  }
}

func (X *console) CircleFullInv (x, y int, r uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.circ (x, y, r, true, X.PointInv)
  }
}

func (X *console) arc (x, y int, r uint, a, b float64, filled bool, f pointFunc) {
  if filled { ker.Panic ("filled arcs not yet implemented") }
// lousy implementation, but better than nothing
  a0, b0, r0, db := a / 180 * math.Pi, b / 180 * math.Pi, float64(r), 1.0 / 180 * math.Pi
  a1 := a0; if b0 > 0 { a1 += b0 } else { a0 += b0 }
  var x1, y1 []int
  for alpha := a0; alpha < a1; alpha += db {
    x1, y1 = append (x1, x + int(r0 * math.Cos(alpha))), append(y1, y - int(r0 * math.Sin(alpha)))
  }
  x1, y1 = append (x1, x + int(r0 * math.Cos(a1))), append(y1, y - int(r0 * math.Sin(a1)))
  for i := 1; i < len(x1); i+= 1 {
    X.line (x1[i-1], y1[i-1], x1[i], y1[i], f)
  }
}

func (X *console) Arc (x, y int, r uint, a, b float64) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.arc (x, y, r, a, b, false, X.Point)
  }
}

func (X *console) ArcInv (x, y int, r uint, a, b float64) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.arc (x, y, r, a, b, false, X.PointInv)
  }
}

func (X *console) ArcFull (x, y int, r uint, a, b float64) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.arc (x, y, r, a, b, true, X.Point)
  }
}

func (X *console) ArcFullInv (x, y int, r uint, a, b float64) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.arc (x, y, r, a, b, true, X.PointInv)
  }
}

func (X *console) OnCircle (x, y int, r uint, a, b int, t uint) bool {
//  if ! between (x - int(r), x + int(r), a) { return false }
/*
  if r == 0 { return a == x && b == y }
  z = a * a + b * b
  if z > r * r { z = z - r * r } else { z = r * r - z }
*/
  X.xx_, X.yy_, X.tt_, X.incident = a, b, int(t * t), false
  X.circ (x, y, r, false, X.onPoint)
  return X.incident
}

func (X *console) ell (x, y int, a, b uint, filled bool, f pointFunc) {
  if ! X.iok (x, y) { return }
  if a == b {
    X.circ (x, y, a, filled, f)
    return
  }
  if a == 0 {
    if b == 0 {
      f (x, y)
    } else {
      X.vertical (x, y - int(b), y + int(b), f)
    }
    return
  } else {
    if b == 0 {
      X.horizontal (x - int(a), y, x + int(a), f)
      return
    }
  }
  a1, b1 := 2 * a * a, 2 * b * b
  i := int (a * b * b)
  x2, y2 := int(2 * a * b * b), 0
  xi, x1 := x - int(a), x + int(a)
  yi, y1 := y, y
  var xl int
  if xi < 0 {
    xl = 0
  } else {
    xl = xi
  }
  if filled {
    X.horizontal (xl, y, x1, f)
  } else {
    f (xl, y)
    f (int(x1), y)
  }
  var yo int
  if a == 0 {
    if y < int(b) {
      yo = 0
    } else {
      yo = y - int(b)
    }
    X.vertical (xi, yo, y + int(b), f)
    return
  }
  for { // a > uint(0) {
    if i > 0 {
      yi--
      y1++
      y2 += int(a1)
      i -= int(y2)
    }
    if i <= 0 {
      xi++
      x1--
      x2 -= int(b1)
      i += int(x2)
      a--
    }
    if xi < 0 {
      xl = 0
    } else {
      xl = xi
    }
    if yi < 0 {
      yo = 0
    } else {
      yo = yi
    }
    var xr int
    if x1 < int(X.wd) {
      xr = int(x1)
    } else {
      xr = int(X.wd) - 1
    }
    var yu int
    if y1 < int(X.ht) {
      yu = int(y1)
    } else {
      yu = int(X.ht) - 1
    }
    if filled {
      X.horizontal (xl, yo, xr, f)
      X.horizontal (xl, yu, xr, f)
    } else {
      f (xl, yo)
      f (xr, yo)
      f (xl, yu)
      f (xr, yu)
    }
    if a == uint(0) {
      break
    }
  }
}

func (X *console) Ellipse (x, y int, a, b uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= a && uint(y) >= b {
    X.ell (x, y, a, b, false, X.Point)
  }
}

func (X *console) EllipseInv (x, y int, a, b uint) {
  if ! X.iok (x, y) { return }
   if uint(x) >= a && uint(y) >= b {
    X.ell (x, y, a, b, false, X.PointInv)
  }
}

func (X *console) EllipseFull (x, y int, a, b uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= a && uint(y) >= b {
    X.ell (x, y, a, b, true, X.Point)
  }
}


func (X *console) EllipseFullInv (x, y int, a, b uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= a && uint(y) >= b {
    X.ell (x, y, a, b, true, X.PointInv)
  }
}

func (X *console) OnEllipse (x, y int, a, b uint, A, B int, t uint) bool {
  if ! X.iok (x, y) { return false }
  X.xx_, X.yy_, X.tt_, X.incident = A, B, int(t * t), false
  X.ell (x, y, a, b, false, X.onPoint)
  return X.incident
}

func (X *console) curve (xs, ys []int, f pointFunc) {
  m := len (xs)
  if m == 0 || m != len (ys) {
panic ("curve: wrong m")
    return
  }
  n := ker.ArcLen (xs, ys)
  xs1, ys1 := make ([]int, n), make ([]int, n)
  for i := uint(0); i < n; i++ {
    xs1[i], ys1[i] = ker.Bezier (xs, ys, uint(m), n, i)
  }
  f (xs[0], ys[0])
  for i := 0; i < len(xs1); i++ {
    f (xs1[i], ys1[i])
  }
}

func (X *console) Curve (xs, ys []int) {
  X.curve (xs, ys, X.Point)
}

func (X *console) CurveInv (xs, ys []int) {
  X.curve (xs, ys, X.PointInv)
}

func (X *console) OnCurve (xs, ys []int, a, b int, t uint) bool {
  if ! ok2 (xs, ys) {
panic ("OnCurve: ! ok2")
    return false
  }
  X.xx_, X.yy_, X.tt_, X.incident = a, b, int(t * t), false
  X.curve (xs, ys, X.onPoint)
  return X.incident
}
