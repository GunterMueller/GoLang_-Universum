package ker

// (c) Christian Maurer   v. 130326 - license see µu.go

import
  . "runtime"

func init() {
  GOMAXPROCS (NumCPU())
}
