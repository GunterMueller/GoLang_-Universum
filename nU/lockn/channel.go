package lockn

// (c) Christian Maurer   v. 190325 - license see nU.go

// >>> Implementation with message passing

type
  channel struct {
               c chan bool
                 }

func newChannel (n uint) LockerN {
  x := new(channel)
  x.c = make(chan bool, 1)
  x.c <- true
  return x
}

func (x *channel) Lock (p uint) {
  <-x.c
}

func (x *channel) Unlock (p uint) {
  x.c <- true
}
