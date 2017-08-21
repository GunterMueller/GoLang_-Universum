package lockp

// (c) murus.org  v. 170624 - license see murus.org

// Szymanski, B. K: A Simple Solution to Lamport's Concurrent Programming Problem with Linear Wait.
// In: Lenfant, J. (ed.): ICS '88, New York. ACM Press (1988) 621-626

type
  szymanski struct {
                   uint "number of processes"
            intent,
   doorIn, doorOut []bool
                   }

func newSz (n uint) LockerP {
  if n < 2 { return nil }
  x := new(szymanski)
  x.uint = n
  x.intent = make([]bool, n)
  x.doorIn, x.doorOut = make([]bool, n), make([]bool, n)
  return x
}

func (x *szymanski) Lock (i uint) { // < x.uint
// P10: flag[i] = 1
  x.intent[i] = true
// P11: wait until for all j: flag[j] <= 2
  j := uint(0)
  for j < x.uint {
    if x.intent[j] && x.doorIn[j] {
      j = 0
    } else {
      j++
    }
  }
// P20: flag[i] = 3
  x.doorIn[i] = true
P21: // if exists j: flag[j] == 1 {
  j = 0
  for (!x.intent[j] || x.doorIn[j]) && j + 1 < x.uint {
    j++
  }
  if x.intent[j] && !x.doorIn[j] {
//  flag[i] = 2
    x.intent[i] = false
// P22: wait until exists j: flag[j] == 4 }
/*
    for !x.doorOut[j] &&
        (j + 1 < N || f = x.intent[j] && !doorIn[j] || f && j > 0) {
*/
    for !x.doorOut[j] {
      f := x.intent[j]
      if j + 1 >= x.uint && (!f || x.doorIn[j]) && (!f || j == 0) {
        break
      }

      if j + 1 == x.uint {
        j = 0
      } else {
        j++
      }
    }
    x.intent[i] = true
    if !x.doorOut[i] {
      goto P21
    }
  }
// P30: flag[i] = 4
  x.doorOut[i] = true
// P31: wait until for all j: j >= i || flag[j] <= 1
  j = 0
  for j < i {
    if x.doorIn[j] {
      j = 0
    } else {
      j++
    }
  }
}

func (x *szymanski) Unlock (i uint) {
// E0: wait until for all j: j <= i || flag[j] <= 1 || flag[j] == 4
  j := i + 1
  for j < x.uint {
    if x.doorIn[j] && !x.doorOut[j] {
      j = i + 1
    } else {
      j++
    }
  }
// E1: flag[i] = 0
  x.intent[i], x.doorIn[i], x.doorOut[i] = false, false, false
}
