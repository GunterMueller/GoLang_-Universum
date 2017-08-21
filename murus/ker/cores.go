package ker

// (c) murus.org  v. 130326 - license see murus.go

import
  . "runtime"

func init() {
  GOMAXPROCS (NumCPU())
}
