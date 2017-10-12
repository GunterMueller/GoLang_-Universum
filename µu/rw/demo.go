package rw

// (c) Christian Maurer   v. 170411 - license see µu.go

import
  . "µu/obj"
var
  writeR, writeW = Ignore, Ignore

func InstallDemo (r, w Op) {
  writeR, writeW = r, w
}
