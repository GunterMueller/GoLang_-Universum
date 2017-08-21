package rw

// (c) murus.org  v. 170411 - license see murus.go

// >>> Solution with Go's sync.RWMutex,
//     most likely the most efficient solution

import
  . "sync"
type
  readerWriter struct {
                      RWMutex
                      }

func new_() ReaderWriter {
  return new (readerWriter)
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

func (x *readerWriter) Fin() {
}
