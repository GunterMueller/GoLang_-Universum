package scr

// (c) Christian Maurer   v. 180421 - license see µU.go

import
  . "µU/linewd"

func (X *screen) ActLinewidth() Linewidth {
  if underX {
    return X.XWindow.ActLinewidth()
  }
  return X.Console.ActLinewidth()
}

func (X *screen) SetLinewidth (w Linewidth) {
  if underX {
    X.XWindow.SetLinewidth (w)
  } else {
    X.Console.SetLinewidth (w)
  }
}

func (X *screen) Point (x, y int) {
  if underX {
    X.XWindow.Point (x, y)
  } else {
    X.Console.Point (x, y)
  }
}

func (X *screen) PointInv (x, y int) {
  if underX {
    X.XWindow.PointInv (x, y)
  } else {
    X.Console.PointInv (x, y)
  }
}

func (X *screen) Points (xs, ys []int) {
  if underX {
    X.XWindow.Points (xs, ys)
  } else {
    X.Console.Points (xs, ys)
  }
}

func (X *screen) PointsInv (xs, ys []int) {
  if underX {
    X.XWindow.PointsInv (xs, ys)
  } else {
    X.Console.PointsInv (xs, ys)
  }
}

func (X *screen) OnPoint (x, y, a, b int, d uint) bool {
  if underX {
    return X.XWindow.OnPoint (x, y, a, b, d)
  } else {
    return X.Console.OnPoint (x, y, a, b, d)
  }
}

func (X *screen) Line (x, y, x1, y1 int) {
  if underX {
    X.XWindow.Line (x, y, x1, y1)
  } else {
    X.Console.Line (x, y, x1, y1)
  }
}

func (X *screen) LineInv (x, y, x1, y1 int) {
  if underX {
    X.XWindow.LineInv (x, y, x1, y1)
  } else {
    X.Console.LineInv (x, y, x1, y1)
  }
}

func (X *screen) OnLine (x, y, x1, y1, a, b int, t uint) bool {
  if underX {
    return X.XWindow.OnLine (x, y, x1, y1, a, b, t)
  }
  return X.Console.OnLine (x, y, x1, y1, a, b, t)
}

func (X *screen) Lines (xs, ys, xs1, ys1 []int) {
  if underX {
    X.XWindow.Lines (xs, ys, xs1, ys1)
  } else {
    X.Console.Lines (xs, ys, xs1, ys1)
  }
}

func (X *screen) LinesInv (xs, ys, xs1, ys1 []int) {
  if underX {
    X.XWindow.LinesInv (xs, ys, xs1, ys1)
  } else {
    X.Console.LinesInv (xs, ys, xs1, ys1)
  }
}

func (X *screen) OnLines (xs, ys, xs1, ys1 []int, a, b int, t uint) bool {
  if underX {
    return X.XWindow.OnLines (xs, ys, xs1, ys1, a, b, t)
  }
  return X.Console.OnLines (xs, ys, xs1, ys1, a, b, t)
}

func (X *screen) Segments (xs, ys []int) {
  if underX {
    X.XWindow.Segments (xs, ys)
  } else {
    X.Console.Segments (xs, ys)
  }
}

func (X *screen) SegmentsInv (xs, ys []int) {
  if underX {
    X.XWindow.SegmentsInv (xs, ys)
  } else {
    X.Console.SegmentsInv (xs, ys)
  }
}

func (X *screen) OnSegments (xs, ys []int, a, b int, t uint) bool {
  if underX {
    return X.XWindow.OnSegments (xs, ys, a, b, t)
  }
  return X.Console.OnSegments (xs, ys, a, b, t)
}

func (X *screen) InfLine (x, y, x1, y1 int) {
  if underX {
    X.XWindow.InfLine (x, y, x1, y1)
  } else {
    X.Console.InfLine (x, y, x1, y1)
  }
}

func (X *screen) InfLineInv (x, y, x1, y1 int) {
  if underX {
    X.XWindow.InfLineInv (x, y, x1, y1)
  } else {
    X.Console.InfLineInv (x, y, x1, y1)
  }
}

func (X *screen) OnInfLine (x, y, x1, y1, a, b int, t uint) bool {
  if underX {
    return X.XWindow.OnInfLine (x, y, x1, y1, a, b, t)
  }
  return X.Console.OnInfLine (x, y, x1, y1, a, b, t)
}

func (X *screen) Triangle (x, y, x1, y1, x2, y2 int) {
  if underX {
    X.XWindow.Triangle (x, y, x1, y1, x2, y2)
  } else {
    X.Console.Triangle (x, y, x1, y1, x2, y2)
  }
}

func (X *screen) TriangleInv (x, y, x1, y1, x2, y2 int) {
  if underX {
    X.XWindow.TriangleInv (x, y, x1, y1, x2, y2)
  } else {
    X.Console.TriangleInv (x, y, x1, y1, x2, y2)
  }
}

func (X *screen) TriangleFull (x, y, x1, y1, x2, y2 int) {
  if underX {
    X.XWindow.TriangleFull (x, y, x1, y1, x2, y2)
  } else {
    X.Console.TriangleFull (x, y, x1, y1, x2, y2)
  }
}

func (X *screen) TriangleFullInv (x, y, x1, y1, x2, y2 int) {
  if underX {
    X.XWindow.TriangleFullInv (x, y, x1, y1, x2, y2)
  } else {
    X.Console.TriangleFullInv (x, y, x1, y1, x2, y2)
  }
}

func (X *screen) Rectangle (x, y, x1, y1 int) {
  if underX {
    X.XWindow.Rectangle (x, y, x1, y1)
  } else {
    X.Console.Rectangle (x, y, x1, y1)
  }
}

func (X *screen) RectangleInv (x, y, x1, y1 int) {
  if underX {
    X.XWindow.RectangleInv (x, y, x1, y1)
  } else {
    X.Console.RectangleInv (x, y, x1, y1)
  }
}

func (X *screen) RectangleFull (x, y, x1, y1 int) {
  if underX {
    X.XWindow.RectangleFull (x, y, x1, y1)
  } else {
    X.Console.RectangleFull (x, y, x1, y1)
  }
}

func (X *screen) RectangleFullInv (x, y, x1, y1 int) {
  if underX {
    X.XWindow.RectangleFullInv (x, y, x1, y1)
  } else {
    X.Console.RectangleFullInv (x, y, x1, y1)
  }
}

func (X *screen) OnRectangle (x, y, x1, y1, a, b int, t uint) bool {
  if underX {
    return X.XWindow.OnRectangle (x, y, x1, y1, a, b, t)
  }
  return X.Console.OnRectangle (x, y, x1, y1, a, b, t)
}


func (X *screen) InRectangle (x, y, x1, y1, a, b int, t uint) bool {
  if underX {
    return X.XWindow.InRectangle (x, y1, x1, y1, a, b, t)
  }
  return X.Console.InRectangle (x, y, x1, y1, a, b, t)
}

func (X *screen) Polygon (xs, ys []int) {
  if underX {
    X.XWindow.Polygon (xs, ys)
  } else {
    X.Console.Polygon (xs, ys)
  }
}

func (X *screen) PolygonInv (xs, ys []int) {
  if underX {
    X.XWindow.PolygonInv (xs, ys)
  } else {
    X.Console.PolygonInv (xs, ys)
  }
}

func (X *screen) PolygonFull (xs, ys []int) {
  if underX {
    X.XWindow.PolygonFull (xs, ys)
  } else {
    X.Console.PolygonFull (xs, ys)
  }
}

func (X *screen) PolygonFullInv (xs, ys []int) {
  if underX {
    X.XWindow.PolygonFullInv (xs, ys)
  } else {
    X.Console.PolygonFullInv (xs, ys)
  }
}

func (X *screen) OnPolygon (xs, ys []int, a, b int, t uint) bool {
  if underX {
    return X.XWindow.OnPolygon (xs, ys, a, b, t)
  }
  return X.Console.OnPolygon (xs, ys, a, b, t)
}

func (X *screen) Circle (x, y int, r uint) {
  if underX {
    X.XWindow.Circle (x, y, r)
  } else {
    X.Console.Circle (x, y, r)
  }
}

func (X *screen) CircleInv (x, y int, r uint) {
  if underX {
    X.XWindow.CircleInv (x, y, r)
  } else {
    X.Console.CircleInv (x, y, r)
  }
}

func (X *screen) CircleFull (x, y int, r uint) {
  if underX {
    X.XWindow.CircleFull (x, y, r)
  } else {
    X.Console.CircleFull (x, y, r)
  }
}

func (X *screen) CircleFullInv (x, y int, r uint) {
  if underX {
    X.XWindow.CircleFullInv (x, y, r)
  } else {
    X.Console.CircleFullInv (x, y, r)
  }
}

func (X *screen) OnCircle (x, y int, r uint, a, b int, t uint) bool {
  if underX {
    return X.XWindow.OnCircle (x, y, r,a, b, t)
  }
  return X.Console.OnCircle (x, y, r, a, b, t)
}

func (X *screen) Arc (x, y int, r uint, a, b float64) {
  if underX {
    X.XWindow.Arc (x, y, r, a, b)
  } else {
    X.Console.Arc (x, y, r, a, b)
  }
}

func (X *screen) ArcInv (x, y int, r uint, a, b float64) {
  if underX {
    X.XWindow.ArcInv (x, y, r, a, b)
  } else {
    X.Console.ArcInv (x, y, r, a, b)
  }
}

func (X *screen) ArcFull (x, y int, r uint, a, b float64) {
  if underX {
    X.XWindow.ArcFull (x, y, r, a, b)
  } else {
    X.Console.ArcFull (x, y, r, a, b)
  }
}

func (X *screen) ArcFullInv (x, y int, r uint, a, b float64) {
  if underX {
    X.XWindow.ArcFullInv (x, y, r, a, b)
  } else {
    X.Console.ArcFullInv (x, y, r, a, b)
  }
}

func (X *screen) Ellipse (x, y int, a, b uint) {
  if underX {
    X.XWindow.Ellipse (x, y, a, b)
  } else {
    X.Console.Ellipse (x, y, a, b)
  }
}

func (X *screen) EllipseInv (x, y int, a, b uint) {
  if underX {
    X.XWindow.EllipseInv (x, y, a, b)
  } else {
    X.Console.EllipseInv (x, y, a, b)
  }
}

func (X *screen) EllipseFull (x, y int, a, b uint) {
  if underX {
    X.XWindow.EllipseFull (x, y, a, b)
  } else {
    X.Console.EllipseFull (x, y, a, b)
  }
}

func (X *screen) EllipseFullInv (x, y int, a, b uint) {
  if underX {
    X.XWindow.EllipseFullInv (x, y, a, b)
  } else {
    X.Console.EllipseFullInv (x, y, a, b)
  }
}

func (X *screen) OnEllipse (x, y int, a, b uint, A, B int, t uint) bool {
  if underX {
    return X.XWindow.OnEllipse (x, y, a, b, A, B, t)
  }
  return X.Console.OnEllipse (x, y, a, b, A, B, t)
}

func (X *screen) Curve (xs, ys []int) {
  if underX {
    X.XWindow.Curve (xs, ys)
  } else {
    X.Console.Curve (xs, ys)
  }
}

func (X *screen) CurveInv (xs, ys []int) {
  if underX {
    X.XWindow.CurveInv (xs, ys)
  } else {
    X.Console.CurveInv (xs, ys)
  }
}

func (X *screen) OnCurve (xs, ys []int, a, b int, t uint) bool {
  if underX {
    return X.XWindow.OnCurve (xs, ys, a, b, t)
  }
  return X.Console.OnCurve (xs, ys, a, b, t)
}
