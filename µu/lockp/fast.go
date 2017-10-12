package lockp

// (c) Christian Maurer   v. 170807 - license see µu.go

// >>> fast algorithm
// see Lamport, L.: A fast mutual exclusion algorithm. ACM TOCS 5 (1987), 1-11

type
  lockerPfast struct {
          nProcesses uint
          interested []bool
         gate, gate1 uint
                     }

func newF (n uint) LockerP {
  if n < 2 { return nil }
  x := new(lockerPfast)
  x.nProcesses = n
  x.interested = make ([]bool, n)
  x.gate, x.gate1 = x.nProcesses, x.nProcesses
  return x
}

func (x *lockerPfast) Lock (p uint) {
  if p >= x.nProcesses { return }
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
    for i:= uint(0); i < x.nProcesses; i++ {
      if i != p {
        for x.interested[i] { /* Null() */ }
      }
    }
    if x.gate1 != p {
      goto p1
    } else {
      x.interested[p] = true
    }
  }
}

func (x *lockerPfast) Unlock (p uint) {
  if p >= x.nProcesses { return }
  x.gate1 = x.nProcesses
  x.interested[p] = false
}
