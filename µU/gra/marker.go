package gra

// (c) Christian Maurer   v. 171112 - license see µU.go

func (x *graph) Mark (m bool) {
  if x.local == nil {
    return
  }
  x.local.marked = m
}

func (x *graph) Marked() bool {
  if x.vAnchor == x.vAnchor.nextV {
    return false
  }
  return x.local.marked
}
