package rw

import (
  "sync"
  "µU/sem"
)
type
  kanglee struct {
  nR, nW, bR, bW uint
                 sync.Mutex
            r, w sem.Semaphore
}

func newKL() ReaderWriter {
  x := new(kanglee)
  x.r = sem.New(1)
  x.w = sem.New(1)
  return x
}

func (x *kanglee) ReaderIn() {
  x.Mutex.Lock()
  if x.nW + x.bW == 0 {
    x.nR++
    x.Mutex.Unlock()
  } else {
    x.bR++
    x.Mutex.Unlock()
    x.r.P()
  }
}

func (x *kanglee) ReaderOut() {
  x.Mutex.Lock()
  x.nR--
  if x.nR == 0 && x.bW > 0 {
    x.nW++
    x.bW--
    x.w.V()
  }
  x.Mutex.Unlock()
}

func (x *kanglee) WriterIn() {
  x.Mutex.Lock()
  if x.nR + x.nW + x.bR + x.bW == 0 {
    x.nW++
    x.Mutex.Unlock()
  } else {
    x.bW++
    x.Mutex.Unlock()
    x.w.P()
  }
}

func (x *kanglee) WriterOut() {
  x.Mutex.Lock()
  x.nW--
  if x.bR > 0 {
    for x.bR > 0 {
      x.bR--
      x.nR++
      x.r.V()
    }
  } else {
    if x.bW > 0 {
      x.nW++
      x.bW--
      x.w.V()
    }
  }
  x.Mutex.Unlock()
}

func (x *kanglee) Fin() {
}

/*
// reader priority

func newKL() ReaderWriter {
  x := new(kanglee)
  x.r = sem.New(0)
  x.w = sem.New(0)
  return x
}

func ReaderIn() {
  x.Mutex.Lock()
  if x.nW == 0 {
    x.nR++
    x.Mutex.Unlock()
  } else {
    x.bR++
    x.Mutex.Unlock()
    x.r.P()
  }
}

func ReaderOut() {
  x.Mutex.Lock()
  x.nR--
  if x.nR == 0 && x.bW > 0 {
    x.bW--
    x.nW++
    x.w.V()
  }
  x.Mutex.Unlock()
}

func WriterIn() {
  x.Mutex.Lock()
    if x.nR + x.nW + x.bR + x.bW == 0 {
    x.nW++
    x.Mutex.Unlock()
  } else {
    x.bW++
    x.Mutex.Unlock()
    x.w.P()
  }
}

func WriterOut() {
  x.Mutex.Lock()
  x.nW--
  if x.bR > 0 {
    for x.bR > 0 {
      x.bR--
      x.nR++
      x.r.V()
    }
  } else {
    if x.bW > 0 {
      x.bW--
      x.nW++
      x.w.V()
    }
  }
  x.Mutex.Unlock()
}
*/

/*
// writer priority

func newKL() ReaderWriter {
  x := new(kanglee)
  x.r = sem.New(0)
  x.w = sem.New(0)
  return x
}

func ReaderIn() {
  x.Mutex.Lock()
  if x.nW + x.bW == 0 {
    x.nR++
    x.Mutex.Unlock()
  } else {
    x.bR++
    x.Mutex.Unlock()
    x.r.P()
  }
}

func ReaderOut() {
  x.Mutex.Lock()
  x.nR--
  if x.nR == 0 && x.bW > 0 {
    x.bW--
    x.nW++
    x.w.V()
  }
  x.Mutex.Unlock()
}

func WriterIn() {
  x.Mutex.Lock()
  if x.nR + x.nW + x.bR + x.bW == 0 {
    x.nW++
    x.Mutex.Unlock()
  } else {
    x.bW++
    x.Mutex.Unlock()
    x.w.P()
  }
}

func WriterOut() {
  x.Mutex.Lock()
  x.nW--
  if x.bW > 0 {
    x.nW++
    x.bW--
    x.w.V()
  } else {
    for x.bR > 0 {
      x.bR--
      x.nR++
      x.r.V()
    }
  }
  x.Mutex.Unlock()
}
*/
