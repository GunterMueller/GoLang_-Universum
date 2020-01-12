package ker

// (c) Christian Maurer   v. 190811 - license see µU.go

import
  "runtime"

func Is64bit() bool {
  switch runtime.GOARCH {
  case "386":
    return false
  case "amd64":
    return true
  }
  Panic ("$GOARCH not treated")
  return false
}
