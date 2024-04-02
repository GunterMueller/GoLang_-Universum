package ker

// (c) Christian Maurer   v. 240401 - license see µU.go

import
  "math"
const (
  Dot = "."
  Mu = "µU"
  DotMu = Dot + Mu

  MaxShortNat = uint(math.MaxUint16)
  MaxNat = uint(math.MaxUint64)
  MaxInt = math.MaxInt64
  MinInt = math.MinInt64

  PointsPerInch = 72
  MillimetersPerInch = 25.4
  PointsPerMillimeter = float64(PointsPerInch) / MillimetersPerInch
                        // 1 mm = 2.834645669291338582677165354330708661417322 pt
  PointsPerCentimeter = 10 * PointsPerMillimeter

  A3wd = 297 // mm
  A3ht = 420 // mm
  A3wdPt = 842 // A3wd * PointsPerMillimeter // pt
  A3htPt = 1191 // A3ht * PointsPerMillimeter // pt

  A4wd = 210 // mm
  A4ht = 297 // mm
  A4wdPt = 595 // A4wd * PointsPerMillimeter // pt
  A4htPt = 842 // A4ht * PointsPerMillimeter // pt

  A5wd = 148 // mm
  A5ht = 210 // mm
  A5wdPt = 421 // A5wd * PointsPerMillimeter // pt
  A5htPt = 595 // A5ht * PointsPerMillimeter // pt
)
