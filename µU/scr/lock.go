package scr

// (c) Christian Maurer   v. 230306 - license see µU.go

import
  "sync"
var
  mutex, mutex1, mutex2 sync.Mutex

func lock() {
  mutex.Lock()
}

func unlock() {
  mutex.Unlock()
}

func lock1() {
  mutex1.Lock()
}

func unlock1() {
  mutex1.Unlock()
}

func lock2() {
  mutex2.Lock()
}

func unlock2() {
  mutex2.Unlock()
}
