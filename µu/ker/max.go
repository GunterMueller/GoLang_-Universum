package ker

// (c) Christian Maurer   v. 150915 - license see µu.go

import (
  "unsafe"
  "math"
)

func MaxNat() uint {
  if unsafe.Sizeof(uint(0)) == 32 {
    return math.MaxUint32
  }
  return math.MaxUint64
}

const
  MaxShortNat = uint(math.MaxUint16)

func MaxInt() int {
  if unsafe.Sizeof(int(0)) == 32 {
    return math.MaxInt32
  }
  return math.MaxInt64
}

func MinInt() int {
  if unsafe.Sizeof(int(0)) == 32 {
    return math.MinInt32
  }
  return math.MinInt64
}
