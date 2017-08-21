package rw

// (c) murus.org  v. 170411 - license see murus.go

import
  . "murus/obj"
var
  writeR, writeW = Ignore, Ignore

func InstallDemo (r, w Op) {
  writeR, writeW = r, w
}
