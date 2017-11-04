package mbuf

// (c) Christian Maurer   v. 171104 - license see µU.go

// >>> implementation with synchronous message passing ans selective waiting

import
  . "µU/obj"
type
  channel1 struct {
                  Any
       cIns, cGet chan Any
                  }

func newCh1 (a Any, n uint) MBuffer {
  if a == nil || n == 0 { return nil }
  x := new(channel1)
  x.Any = Clone (a)
  x.cIns, x.cGet = make(chan Any), make(chan Any)
  go func() {
    buffer := make([]Any, n)
    var count, in, out uint
    for {
      if count == 0 {
        buffer[in] = <-x.cIns
        in = (in + 1) % n; count = 1
      } else if count == n {
        x.cGet <- buffer[out]
        out = (out + 1) % n; count = n - 1
      } else { // 0 < count < n
        select {
        case buffer[in] = <-x.cIns:
          in = (in + 1) % n; count++
        case x.cGet <- buffer[out]:
          out = (out + 1) % n; count--
        }
      }
    }
  }()
  return x
}

func (x *channel1) Ins (a Any) { x.cIns <- a }
func (x *channel1) Get() Any { return <-x.cGet }
