package dgra

// (c) Christian Maurer   v. 171220 - license see µU.go

import
  . "µU/obj"
const (
  DISCOVER = uint(iota)
  RETURN
)

func (x *distributedGraph) helaryRaynal (o Op) {
  x.connect (nil)
  defer x.fin()
  x.Op = o
  if x.me == x.root {
    x.parent = x.root
    us := []uint {DISCOVER, x.me}
x.log2("send to", x.nr[0], "len", Codelen(us))
    x.ch[0].Send (us)
  }
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
//    loop:
      for {
        r := x.ch[j].Recv()
if r == nil { panic("Scheiße") }
        bs := r.(Stream)
x.log2("recv from", x.nr[j], "len", uint(len(bs)))
        us := Decode (UintStream{}, bs).(UintStream)
        v := x.n
        k := uint(0)
        for _, u := range us {
          for i := uint(0); i < x.n; i++ {
            if u == x.nr[i] {
              v--
              break
            } else {
              k = i
            }
          }
        }
        if us[0] == DISCOVER {
          x.parent = x.nr[j]
          us = append(us, x.me)
          if v == 0 {
            us[0] = RETURN
x.log2("send RETURN to", x.nr[j], "len", Codelen(us))
            x.ch[j].Send (us)
          } else {
x.log2("send DISCOVER to", x.nr[k], "len", Codelen(us))
            x.ch[k].Send (us) // DISCOVER
          }
        } else { // us[0] == RETURN
          if v == 0 {
            if x.parent == x.me {
              done <- 0
              x.log0 ("algorithm terminated")
              return
            } else {
x.log2("send RETURN to", x.nr[x.channel(x.parent)], "len", Codelen(us))
              x.ch[x.channel(x.parent)].Send (us) // RETURN
            }
          } else {
            us[0] = DISCOVER
x.log2("send DISCOVER to", x.nr[k], "len", Codelen(us))
            x.ch[k].Send (us)
          }
        }
      }
      done <- 0
    }(i)
  }
  x.Op(x.me)
  for i := uint(0); i < x.n; i++ {
    <-done
  }
}
