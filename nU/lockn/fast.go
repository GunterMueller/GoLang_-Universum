package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

// >>> fast algorithm
// see Lamport, L.: A fast mutual exclusion algorithm. ACM TOCS 5 (1987), 1-11

type lockerNfast struct {
  uint "Anzahl der beteiligten Prozesse"
  interested []bool
  gate, gate1 uint
}

func newF (n uint) LockerN {
  if n < 2 { return nil }
  x := new(lockerNfast)
  x.uint = n
  x.interested = make([]bool, n)
  x.gate, x.gate1 = x.uint, x.uint
  return x
}

func (x *lockerNfast) Lock (p uint) {
  if p >= x.uint { return }
p1:
  x.gate = p
  x.interested[p] = true
  if x.gate1 != p {
    x.interested[p] = false
    goto p1
  }
  x.gate1 = p
  if x.gate != p {
    x.interested[p] = false
    for i:= uint(0); i < x.uint; i++ {
      if i != p {
        for x.interested[i] {
          nothing()
        }
      }
    }
    if x.gate1 != p {
      goto p1
    } else {
      x.interested[p] = true
    }
  }
}

func (x *lockerNfast) Unlock (p uint) {
  if p >= x.uint { return }
  x.gate1 = x.uint
  x.interested[p] = false
}
