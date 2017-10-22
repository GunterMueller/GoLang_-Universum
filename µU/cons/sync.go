package cons

// (c) Christian Maurer   v. 140203 - license see µU.go

import
  "sync"
var
  writing sync.Mutex

func lock() {
  writing.Lock()
}

func unlock() {
  writing.Unlock()
}
