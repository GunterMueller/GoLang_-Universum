package smok

// (c) Christian Maurer   v. 171102 - license see nU.go

type channel struct {
  ch []chan uint
}

func newCh() Smokers {
  x := new(channel)
  x.ch = make([]chan uint, 4)
  for i := 0; i < 3; i++ {
    x.ch[i] = make(chan uint, 1)
  }
  x.ch[3] = make(chan uint, 1)
  x.ch[3] <- 0
  return x
}

func (x *channel) Agent (u uint) {
  <-x.ch[3]
  x.ch[u] <- 0
}

func (x *channel) SmokerIn (u uint) {
  <-x.ch[u]
}

func (x *channel) SmokerOut() {
  x.ch[3] <- 0
}
