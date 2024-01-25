package lock

// (c) Christian Maurer   v. 161216 - license see nU.go

// >>> Implementation with message passing

type
  channel struct {
               c chan bool
                 }

func newChan() Locker {
  x := new(channel)
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
