package lr

// (c) Christian Maurer   v. 171101 - license see µU.go

// >>> bounded left/right problem

import
  . "µU/obj"
type
  channelBounded struct {
   inL, outL, inR, outR chan Any
                   done chan int
                        }

func newChB (mL, mR uint) LeftRight {
  x := new(channelBounded)
  x.inL, x.outL = make(chan Any), make(chan Any)
  x.inR, x.outR = make(chan Any), make(chan Any)
  x.done = make(chan int)
  go func() {
    var (
      nL, nR,     // active lefties, righties
      tR, tL uint // number of actives within one turn
      bR, bL uint // XXX where are these variables set ? ? ? ? ? ? ? ? ? ? ? ? ? ?
    )
    loop:
    for {
      select {
      case <-x.done:
        break loop
      case <-When (nR == 0 && (bR == 0 || tL < mL), x.inL):
        nL++
        tL++
        tR = 0
      case <-When (nR == 0 && nL > 0, x.outL):
        nL--
      case <-When (nL == 0 && (bL == 0 || tR < mR), x.inR):
        nR++
        tR++
        tL = 0
      case <-When (nL == 0 && nR > 0, x.outR):
        nR--
      }
    }
  }()
  return x
}

func (x *channelBounded) LeftIn() {
  x.inL <- 0
}

func (x *channelBounded) LeftOut() {
  x.outL <- 0
}

func (x *channelBounded) RightIn() {
  x.inR <- 0
}

func (x *channelBounded) RightOut() {
  x.outR <- 0
}

func (x *channelBounded) Fin() {
  x.done <- 0
}
