package lockn

// (c) Christian Maurer   v. 171026 - license see µU.go

// >>> Algorithm of Dijkstra
//     Cooperating Sequential Processes, 0 -> true, 1 -> false

type
  burns struct {
     nProcesses,
       favoured uint
     interested []bool
                }

func newBu (n uint) LockerN {
  x := new(burns)
  x.nProcesses = n
  x.interested = make([]bool, n + 1)
  x.favoured = x.nProcesses
  return x
}

func (x *burns) Lock (i uint) {
  x.interested[i] = true
  x.favoured = i
  var j uint
  for {
    for x.favoured != i {
      x.interested[i] = false
      j = 1
      for j < x.nProcesses && (j == i || ! x.interested[j]) {
        j++
      }
      if j >= x.nProcesses {
        x.interested[i] = true
        x.favoured = i
      }
    }
    j = 1
    for j < x.nProcesses && (j == i || ! x.interested[j]) {
      j++
    }
    if j >= x.nProcesses {
      break
    }
  }
}

func (x *burns) Unlock (i uint) {
  x.interested[i] = false
}
