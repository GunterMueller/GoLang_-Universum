package lock

// (c) Christian Maurer   v. 161216 - license see murus.go

// >>> Implementation with asynchronous message passing
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 

type
  channel struct {
               c chan bool
                 }

func newChan() Locker {
  x := new (channel)
  x.c = make(chan bool, 1)
  x.c <- true
  return x
}

func (x *channel) Lock() {
  <-x.c
}

func (x *channel) Unlock() {
  x.c <- true
}
