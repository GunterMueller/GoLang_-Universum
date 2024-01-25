package mbbuf

// (c) Christian Maurer   v. 220702 - license see nU.go

import
   . "nU/obj"
type
  channel1 struct {
                  any
       cIns, cGet chan any
                  }

func newCh1 (a any, n uint) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new(channel1)
  x.any = Clone (a)
  x.cIns, x.cGet = make(chan any), make(chan any)
  go func() {
    buffer := make([]any, n)
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

func (x *channel1) Ins (a any) { x.cIns <- a }
func (x *channel1) Get() any { return <-x.cGet }
