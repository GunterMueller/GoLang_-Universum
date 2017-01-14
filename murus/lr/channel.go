package lr

// (c) murus.org  v. 140615 - license see murus.go

// >>> left/right problem: implementation with channels
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 183

type
  channel struct {
  lI, lO, rI, rO,
            done chan int
                 }

func NewChannel() LeftRight {
  x:= new (channel)
  x.lI, x.lO = make (chan int), make (chan int)
  x.rI, x.rO = make (chan int), make (chan int)
  x.done = make (chan int)
  go func() {
    var nL, nR int
    for {
//      if _, ok:= <-x.done; ok { break }
      if nL == 0 {
        if nR == 0 {
          select { case <-x.lI:
            nL ++
          case <-x.rI:
            nR ++
          }
        } else { // nR > 0
          select { case <-x.rI:
            nR ++
          case <-x.rO:
            nR --
          }
        }
      } else { // nL > 0
        select { case <-x.lI:
          nL ++
        case <-x.lO:
          nL --
        }
      }
    }
  }()
  return x
}

func (x *channel) LeftIn() {
  x.lI <- 0
}

func (x *channel) LeftOut() {
  x.lO <- 0
}

func (x *channel) RightIn() {
  x.rI <- 0
}

func (x *channel) RightOut() {
  x.rO <- 0
}

func (x *channel) Fin() {
  x.done <- 0
}
