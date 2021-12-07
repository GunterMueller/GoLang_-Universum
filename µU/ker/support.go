package ker

// (c) Christian Maurer   v. 211203 - license see µU.go

import (
  "runtime"
  "µU/env"
)
const
  fail = " not supported by µU"
var (
  under_C, under_X bool
)

func init() {
  goarch, goos := runtime.GOARCH, runtime.GOOS
  switch goarch {
  case "386":
    Panic ("µU does not support 32-bit computers")
  case "amd64":
    switch goos {
    case "linux":
      under_X = env.UnderX()
      under_C = ! under_X
    case "windows":
      under_X = true
      under_C = false
    default:
      Panic (goarch + fail)
    }
  default: // arm64, ios
    Panic (goarch + fail)
  }
}

func underC() bool {
  return under_C
}

func underX() bool {
  return under_X
}
