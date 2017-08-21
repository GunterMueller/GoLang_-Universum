package xwin

// (c) murus.org  v. 170814 - license see murus.go

//  include <X11/X.h>
//  include <X11/Xlib.h>
import
//  "C" // It looks as if there is a BUG in the X kernel:
//      // use of XLock/UnlockDisplay leads to sporadic stops in X-programs.
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
