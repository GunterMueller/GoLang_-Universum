package ker

// (c) Christian Maurer   v. 180901 - license see µU.go

const (
  Dot = "."
  Separator = "/"
  Mu = "µU"
  DotMu = Dot + Mu
//  Zero, One = uint(0), uint(1)
  One = uint(1)

  PointsPerInch = 72
  MillimetersPerInch = 25.4
  PointsPerMillimeter = float64(PointsPerInch) / MillimetersPerInch
                        // 1 mm = 2.834645669291338582677165354330708661417322 pt
  PointsPerCentimeter = 10 * PointsPerMillimeter

  A4wd = 210 // mm
  A4ht = 297 // mm
  A4wdPt = 596 // A4wd * PointsPerMillimeter // pt
  A4htPt = 842 // A4ht * PointsPerMillimeter // pt
  A5wd = 148 // mm
  A5ht = 210 // mm
  A5wdPt = 421
  A5htPt = 596
)
