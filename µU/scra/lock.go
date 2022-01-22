package scr

import
  "sync"
var
  mutex sync.Mutex

func lock() {
//  C.XLockDisplay (dpy)
  mutex.Lock()
}

func unlock() {
//  C.XUnlockDisplay (dpy)
  mutex.Unlock()
}
