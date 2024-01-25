package rw

// (c) Christian Maurer   v. 171125 - license see nU.go

import
  . "sync"
type
  readerWriter struct {
                      RWMutex
                      }

func newG() ReaderWriter {
  return new(readerWriter)
}

func (x *readerWriter) ReaderIn() {
  x.RLock()
}

func (x *readerWriter) ReaderOut() {
  x.RUnlock()
}

func (x *readerWriter) WriterIn() {
  x.Lock()
}

func (x *readerWriter) WriterOut() {
  x.Unlock()
}
