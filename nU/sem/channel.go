package sem

// (c) Christian Maurer   v. 170121 - license see nU.go

type
  channel struct {
            p, v chan int
                 }

func newCh (n uint) Semaphore {
  x := new(channel)
  x.p, x.v = make(chan int), make(chan int)
  go func() {
    val := n
    for {
      if val == 0 {
        <-x.v
        val = 1
      } else { // val > 0
        select {
        case <-x.p:
          val--
        case <-x.v:
          val++
        }
      }
    }
  }()
  return x
}

func (x *channel) P() {
  x.p <- 0
}

func (x *channel) V() {
  x.v <- 0
}
