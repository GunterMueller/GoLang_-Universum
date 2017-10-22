package lockp

type
  nsp_szymanski struct {
                       uint "number of processes"
                intent,
                doorIn,
               doorOut []bool
                       }

func newNSP_Szymanski (n uint) LockerP {
  if n < 2 { return nil }
  x := new(nsp_szymanski)
  x.uint = n
  x.doorIn = make([]bool, n + 1)
  x.doorIn, x.doorOut = make([]bool, n + 1), make([]bool, n + 1)
  return x
}

func (x *nsp_szymanski) Lock (p uint) {
  x.intent[p] = true
  i := uint(0)
  for i < x.uint {
    if x.intent[i] && x.doorIn[i] {
      i = 0
    } else {
      i++
    }
  }
  x.doorIn[p] = true
  i = 0
  for (! x.intent[i] || x.doorIn[i]) && i + 1 < x.uint {
    i++
  }
  if x.intent[i] && ! x.doorIn[i] {
    x.intent[p] = false
    for ! x.doorOut[i] {
      f := x.intent[i]
      if i + 1 >= x.uint && (! f || x.doorIn[i]) && (! f || i == 0) {
        break
      }
      if i == x.uint - 1 {
        i = 0
      } else {
        i++
      }
    }
    x.intent[p] = true
    if ! x.doorOut[p] {
//      goto P21
    }
  }
  x.doorOut[p] = true
  i = p + 1
  for i < x.uint {
    if x.doorIn[i] && ! x.doorOut[i] {
      i = p + 1
    } else {
      i++
    }
  }
  i = 0
  for i < p {
    if x.doorIn[i] {
      i = 0
    } else {
      i++
    }
  }
}

func (x *nsp_szymanski) Unlock(i uint) {
  x.intent[i] = false
  x.doorIn[i], x.doorOut[i] = false, false
}
