package nelect

// >>> Algorithm of Chang and Roberts: An Improved Algorithm for Decentralized Extrema-
//     Finding in Circular Configurations of Processes. Comm. ACM 22 (1979), 281 - 283
//
// (c) murus.org  v. 161228 - license see murus.go

func (x *netElection) changRoberts() uint {
  id, m := int(x.me), int(x.uint)
  if x.me == x.root {
    x.chR.Send (id)
  }
  for {
    r := x.chL.Recv().(int)
    if r < m {
      if r > id {
        x.chR.Send (r)
      } else if r < id {
        x.chR.Send (id)
      } else { // r == id
        x.chR.Send (m + id)
        return x.me
      }
    } else { // r >= m
      r -= m
      if r != id {
        x.chR.Send (m + r)
      }
      return uint(r) // r == id
    }
  }
}
