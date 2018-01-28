package ker

// (c) Christian Maurer   v. 171230 - license see µU.go

import (
  "unsafe"
  "math"
)
var
  is32bit = unsafe.Sizeof(int(0)) == 32

func MaxNat() uint {
  if is32bit {
    return math.MaxUint32
  }
  return math.MaxUint64
}

const
  MaxShortNat = uint(math.MaxUint16)

func MaxInt() int {
  if is32bit {
    return math.MaxInt32
  }
  return math.MaxInt64
}

func MinInt() int {
  if is32bit {
    return math.MinInt32
  }
  return math.MinInt64
}
