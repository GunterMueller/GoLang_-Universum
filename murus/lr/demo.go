package lr

// (c) murus.org  v. 170411 - license see murus.go

import
  . "murus/obj"
var
  writeL, writeR = Ignore, Ignore

func InstallDemo (l, r Op) {
  writeL, writeR = l, r
}
