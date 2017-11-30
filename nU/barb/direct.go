package barb

// (c) Christian Maurer   v. 170731 - license see nU.go

import
  "sync"
type
  barberD struct {
         waiting,
           mutex sync.Mutex
                 int "n"
                 }


func newD() Barber {
  x := new (barberD)
  x.waiting.Lock()
  return x
}

func (x *barberD) Customer() {
  x.mutex.Lock()
  x.int++
  if x.int == 0 {
//    x.mutex.Unlock() // *
    x.waiting.Unlock()
  } else {
    x.mutex.Unlock()
  }
}

func (x *barberD) Barber() {
  x.mutex.Lock()
  x.int--
  if x.int == -1 {
    x.mutex.Unlock()
    x.waiting.Lock()
//    x.mutex.Lock() // *
  }
  x.mutex.Unlock()
}
