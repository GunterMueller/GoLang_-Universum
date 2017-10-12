package pset

// (c) Christian Maurer   v. 170918 - license see µu.go

// TODO nearly all Iterator methods
// BUG  still some subtle bugs in Del
// TODO freelist

import (
  . "µu/obj"
  "µu/ker"
  "µu/str"
  "µu/col"
  "µu/scr"
  "µu/errh"
  "µu/pseq"
  "µu/pset/internal"
)
const (
  N = internal.N
  max = 2 * N
  maxDepth = 31
  suffix = "btr"
)
type (
  persistentSet struct {
                string "name"
            seq pseq.PersistentSequence
            num,
       freelist uint
          empty Object
            tmp,
          newPg,
           nbPg,
         freePg internal.Page
             pg [maxDepth+1]internal.Page // TODO: slices
            idx [maxDepth+1]uint // <= max + 1
             dp uint // <= maxDepth
                }
)

func new_(a Object) PersistentSet {
  x := new(persistentSet)
  x.seq = pseq.New(internal.New(a))
  x.empty = a.Clone().(Object)
  x.tmp = internal.New (a)
  x.newPg = internal.New (a)
  x.nbPg = internal.New (a)
  x.freePg = internal.New (a)
  x.idx [0] = 0
  for t := 0; t <= maxDepth; t++ {
    x.pg[t] = internal.New (a)
  }
  x.pg[0].PutPos (0, 1)
  x.pg[0].PutPos (1, 0) // x.num
  x.pg[0].PutPos (2, 0) // x.freelist
  return x
}

func (x *persistentSet) Fin() {
  x.seq.Fin()
}

func (x *persistentSet) write (p internal.Page, n uint) {
  x.seq.Seek (n)
  if n == 0 { x.pg[1].PutPos (1, x.num) }
  x.seq.Put (p)
}

func (x *persistentSet) read (n uint) internal.Page {
  x.seq.Seek (n)
  return x.seq.Get().(internal.Page)
}

func (x *persistentSet) Name (s string) {
  x.string = s
  x.seq.Name (x.string + "." + suffix)
  if x.seq.Empty() {
    x.num = 0
    x.pg[0].PutPos (0, 1)
    x.pg[0].PutPos (1, 0)
    x.pg[0].PutPos (2, 2)
    x.write (x.pg[0], 0)
    x.write (x.pg[1], 1)
    x.write (x.pg[2], 2)
  } else {
    x.pg [0] = x.read (0)
    x.num = x.pg [0].GetPos (1)
    x.freelist = x.pg [0].GetPos (2)
    x.pg [1] = x.read (x.pg [0].GetPos (0))
    if x.freelist == 0 {
      x.freePg.Clr()
      n := x.seq.Num()
      x.freePg.PutPos (0, n)
      x.freePg.PutPos (max, n)
      x.freePg.PutNum (0)
      x.write (x.freePg, n)
    } else {
      x.freePg = x.read (x.freelist)
    }
  }
  x.Jump (false)
}

func (x *persistentSet) Rename (s string) {
  if str.Empty (s) || s == x.string { return }
  x.string = s
  x.seq.Rename (x.string + "." + suffix)
}

func (x *persistentSet) Empty() bool {
  return x.num == 0
}

func (x *persistentSet) Clr() {
  x.num = 0
  x.seq.Clr()
  for d := 0; d <= maxDepth; d++ {
    x.pg [d].Clr()
  }
  x.pg[0].PutPos (0, 1)
  x.pg[0].PutPos (1, 0)
  x.write (x.pg [0], 0)
  x.write (x.pg [1], 1)
  x.dp = 0
}

func (x *persistentSet) Offc() bool {
  return x.num == 0
}

func (x *persistentSet) Num() uint {
  return x.num
}

func (x *persistentSet) ins (o Object, n uint) {
  x.pg [x.dp].Ins (o, x.idx [x.dp], n)
}

func (x *persistentSet) Ins (a Any) {
  o := a.(Object)
  if x.Ex (a) {
    return
  }
  x.ins (o, 0)
  x.num ++
  x.pg [0].PutPos (1, x.num)
  for {
    if x.pg [x.dp].GetNum() <= max {
      break
    }
// x.num == max + 1  -->  split page:
// leave left part on x.pg [x.dp] and move right part to newPg
    x.newPg.Clr()
    x.newPg.PutNum (N)
    x.pg [x.dp].PutNum (N)
    for i := uint(0); i < N; i++ {
      x.newPg.Put (i, x.pg [x.dp].Get (N + i + 1))
      x.pg [x.dp].Put (i + 1 + N, x.empty)
      x.newPg.PutPos (i, x.pg [x.dp].GetPos (i + 1 + N))
      x.pg [x.dp].PutPos (i + 1 + N, 0)
    }
    x.newPg.PutPos (N, x.pg [x.dp].GetPos (max + 1))
    x.pg [x.dp].PutPos (max + 1, 0)
    n := x.pg [x.dp - 1].GetPos (x.idx [x.dp - 1])
// save middle object in b and overwrite it with x.empty:
    b := x.pg [x.dp].Get (N)
    x.pg [x.dp].Put (N, x.empty)
    x.write (x.pg [x.dp], n)
    n = x.seq.Num()
    x.write (x.newPg, n)
    if x.dp == 1 { // generate new root page
      x.idx [x.dp] = 0
      x.pg [x.dp].PutNum (0)
      x.pg [x.dp].PutPos (0, x.pg [0].GetPos (0))
      x.pg [0].PutPos (0, n + 1) // see above
      x.ins (b, n)
      break
    } else { // lift up former middle object
      x.dp--
      x.ins (b, n)
    }
  }
  n := x.pg [x.dp - 1].GetPos (x.idx [x.dp - 1])
  x.write (x.pg [x.dp], n)
  x.pg [0].PutPos (1, x.num)
  x.write (x.pg [0], 0)
  if ! x.Ex (a) { ker.Oops() }
}

  func rrx (n int) { println ("# ", n) }

func (x *persistentSet) Step (forward bool) {
  if x.num == 0 || x.dp == 0 {
    return
  }
  if forward {
    x.idx [x.dp]++
    if x.pg [x.dp].GetPos (x.idx [x.dp]) == 0 {
      if x.idx [x.dp] < x.pg [x.dp].GetNum() {
        return
      }
      // x.idx [x.dp] == x.pg [x.dp].GetNum()
      for {
        if x.dp == 1 {
          for x.pg [x.dp].GetPos (x.idx [0]) > 0 {
            x.dp++
          }
          x.idx [x.dp]--
          return
        }
        if x.dp == 0 { ker.Panic1("übles Problem hoch", 3) }
        x.dp--
        if x.dp == 0 || x.idx [x.dp] < x.pg [x.dp].GetNum() {
          return
        }
      }
    } else { // x.pg [x.dp].pos [x.idx [x.dp]] > 0
      for {
        n := x.pg [x.dp].GetPos (x.idx [x.dp])
        x.dp++
        x.pg [x.dp] = x.read (n)
        x.idx [x.dp] = 0
        if x.pg [x.dp].GetPos (0) == 0 {
          return
        }
      }
    }
  } else { // backward
rrx (1000 * (1000 * (1000 + int (x.dp)) + int(x.idx [x.dp])) + int(x.pg [x.dp].GetPos (x.idx [x.dp])))
    if x.pg [x.dp].GetPos (x.idx [x.dp]) == 0 {
rrx (1)
      if x.idx [x.dp] == 0 {
rrx (2)
        x.dp--
rrx (3)
        if x.dp == 0 { // walk left downwards
rrx (4)
          x.dp = 1
          x.idx [x.dp] = 0
rrx (5)
        } else { // x.dp > 0
          for {
            if x.idx [x.dp] > 0 {
              x.idx [x.dp]--
rrx (6)
              return
            } else {
rrx (7)
              if x.dp == 0 {
                for x.pg [x.dp].GetPos (0) > 0 {
                  x.dp++
                }
rrx (8)
                return
              } else {
                x.dp--
rrx (9)
              }
            }
          }
        }
      } else { // x.idx [x.dp] > 0
rrx (10)
        x.idx [x.dp]--
rrx (11)
        return
      }
    } else { // x.pg [x.dp].GetPos (x.idx [x.dp]) > 0
      errh.Error2 ("idx", x.idx [x.dp], ">>> Pos", x.pg [x.dp].GetPos (x.idx [x.dp]))
      for {
rrx (12)
        n := x.pg [x.dp].GetPos (x.idx [x.dp])
        x.dp++ // hier ist der Wurm drin
        x.pg [x.dp] = x.read (n)
        if x.pg [x.dp].GetPos (x.pg [x.dp].GetNum()) == 0 {
rrx (1300 + int(x.dp))
          if x.pg [x.dp].GetNum() > 0 {
            x.idx [x.dp] = x.pg [x.dp].GetNum() - 1
          } else {
            ker.Shit() // böses Problem
          }
rrx (14)
          return
        } else {
rrx (15)
          x.idx [x.dp] = x.pg [x.dp].GetNum()
        }
      }
    }
  }
}

func (x *persistentSet) Jump (forward bool) {
  if x.num == 0 {
    return
  }
  x.dp = 1
  for {
    if forward {
      x.idx [x.dp] = x.pg [x.dp].GetNum()
    } else {
      x.idx [x.dp] = 0
    }
    if x.pg [x.dp].GetPos (x.idx [x.dp]) == 0 {
      if forward {
        x.idx [x.dp] = x.pg [x.dp].GetNum() - 1
      } else {
        x.idx [x.dp] = 0
      }
      break
    } else {
      n := x.pg [x.dp].GetPos (x.idx [x.dp])
      x.dp++
      x.pg [x.dp] = x.read (n)
    }
  }
}

func (x *persistentSet) Eoc (forward bool) bool {
  if x.num == 0 { return false }
  if forward {
    for d := uint(1); d <= x.dp; d++ {
      i := x.pg [d].GetNum()
      if d == x.dp && x.pg [d].GetPos (x.idx [d]) == 0 {
        i--
      }
      if i != x.idx [d] {
        return false
      }
    }
    return true
  }
  for d := uint(1); d <= x.dp; d++ {
    if x.idx [d] > 0 {
      return false
    }
  }
  return x.pg [x.dp].GetPos (0) == 0 // x.idx [x.dp] == 0 !
}

func (x *persistentSet) Get() Any {
  if x.dp == 0 {
    ker.Oops()
  }
  if x.num == 0 {
    return x.empty.Clone()
  }
  if x.idx [x.dp] > 100 { ker.Panic1 ("pset.Get", 1000 + x.idx [x.dp]) }
  return x.pg [x.dp].Get (x.idx [x.dp])
}

func (x *persistentSet) Put (a Any) {
  if ! x.Empty() {
    x.Del()
  }
  x.Ins (a)
}

// Pre: d > 1, x.nbPg.GetNum() > N.
func (x *persistentSet) rot (d uint, right bool) {
  i := x.idx [d - 1]
  i1 := i
  if right { // rotation from right neighbour page to x.pg [d] on the left
    i1 ++
    x.pg [d].Put (x.pg [d].GetNum(), x.pg [d - 1].Get (i))
    x.pg [d].IncNum()
    x.pg [d].PutPos (x.pg [d].GetNum(), x.nbPg.GetPos (0))
    x.pg [d - 1].Put (i, x.nbPg.Get (0))
    x.nbPg.RotLeft()
/********************************************************
// with nbPg
    for i := uint(1); i < num; i++ {
      content [i - 1] = content [i]
      pos [i - 1] = pos [i]
    }
    content [num - 1] = empty
    pos [num - 1] = pos [num]
    pos [num] = 0
    num --
********************************************************/
  } else { // rotation from left neighbour page to x.pg [d] on the right
    i1 --
    if x.pg [d].GetNum() == 0 { ker.Oops() }
    x.pg [d].RotRight()
/********************************************************
// with pg [d]
    pos [num + 1] = pos [num]
//  for i := num - 1; i >= 0; i-- { // does not work, because for uint: 0-- == 2^32 - 1  !
    i := num - 1
    for {
      content [i + 1] = content [i]
      pos [i + 1] = pos [i]
      if i == 0 {
        break
      }
      i--
    }
********************************************************/
    x.pg [d].Put (0, x.pg [d - 1].Get (i1))
    x.pg [d].PutPos (0, x.nbPg.GetPos (x.nbPg.GetNum()))
    x.pg [d].IncNum()
    x.pg [d - 1].Put (i1, x.nbPg.Get (x.nbPg.GetNum() - 1))
    x.nbPg.ClrLast()
/********************************************************
func (x *persistentSet) ClrLast() {
// with nbPg
    content [num - 1] = empty
    pos [num - 1] = pos [num]
    pos [num] = 0
    num --
********************************************************/
  }
  x.write (x.pg [d - 1], x.pg [d - 2].GetPos (x.idx [d - 2]))
  x.write (x.pg [d], x.pg [d - 1].GetPos (i))
  x.write (x.nbPg, x.pg [d - 1].GetPos (i1))
}

func (x *persistentSet) join (d uint, right bool) {
  j := x.idx [d - 1]
  j1 := j
  var j0 uint
  if right { // move right neighbour page into x.pg [d]
    j1 ++
    j0 = j1
    n := x.pg [d].GetNum()
    x.pg [d].Put (n, x.pg [d - 1].Get (j))
    x.pg [d].PutNum (N)
    if x.nbPg.GetNum() != N { ker.Oops() }
    n = x.pg [d].GetNum()
    for i := uint(0); i < N; i++ {
      x.pg [d].Put (n + i, x.nbPg.Get (i))
      x.pg [d].PutPos (n + i, x.nbPg.GetPos (i))
    }
    x.pg [d].PutNum (max)
    x.pg [d].PutPos (max, x.nbPg.GetPos (N))
    x.write (x.pg [d], x.pg [d - 1].GetPos (j))
    x.nbPg.Clr()
    x.write (x.nbPg, x.pg [d - 1].GetPos (j1))
  } else { // move x.pg [d] into left neighbour page
    j1 --
    j0 = j
    n := x.nbPg.GetNum()
    if n != N { ker.Oops() }
    x.nbPg.Put (n, x.pg [d - 1].Get (j1))
    n++
    x.nbPg.PutNum (n)
    if x.pg [d].GetNum() != N - 1 { ker.Oops() }
    for i := uint(0); i < N - 1; i++ {
      x.nbPg.Put (n + i, x.pg [d].Get (i))
      x.nbPg.PutPos (n + i, x.pg [d].GetPos (i))
    }
    x.nbPg.PutNum (max)
    x.nbPg.PutPos (x.nbPg.GetNum(), x.pg [d].GetPos (N - 1))
    x.write (x.nbPg, x.pg [d - 1].GetPos (j1))
    x.pg [d].Clr()
    x.write (x.pg [d], x.pg [d - 1].GetPos (j))
  }
  x.pg [d - 1].Join (j0)
/********************************************************
// with pg [d - 1]
  if j0 < num {
    for i := j0; i < num; i++ {
      content [i - 1] = content [i]
      pos [i] = pos [i + 1]
    }
  }
  content [num - 1] = empty
  pos [num] = 0
  num --
********************************************************/
  x.write (x.pg [d - 1], x.pg [d - 2].GetPos (x.idx [d - 2]))
}

func (x *persistentSet) removeUnderflow (d uint) {
  if d == 1 {
    n := x.pg [0].GetPos (0)
    if x.pg [1].GetNum() == 0 {
      x.pg [0].PutPos (0, x.pg [1].GetPos (0))
      x.pg [0].PutPos (1, x.num)
      x.write (x.pg [0], 0)
      x.pg [1].Clr()
      x.write (x.pg [1], n)
    }
    x.write (x.pg [1], n)
    return
  }
  // d > 1
  i := x.idx [d - 1]
  right := i < x.pg [d - 1].GetNum()
  i1 := i
  if right {
    i1 ++ // x.nbPg becomes right neighbour page
  } else { // i == x.pg [d - 1].GetNum()
    i1 -- // x.nbPg becomes left neighbour page
  }
  nn := x.pg [d - 1].GetPos (i1)
  x.nbPg = x.read (nn)
  if x.nbPg.GetNum() > N { // rotation possible
    x.rot (d, right)
  } else { // x.nbPg.GetNum() <= N
    if x.nbPg.GetNum() < N { ker.Todo() } // happens TODO
    x.join (d, right)
    if x.pg [d - 1].GetNum() < N {
      x.removeUnderflow (d - 1)
    }
  }
}

func tst (n uint) {
  println ("Testpunkt ", n)
  return
  errh.Error ("Testpunkt", n)
}

func (x *persistentSet) Del() Any {
  if x.num == 0 {
    return x.empty.Clone()
  } else {
    x.num --
    x.pg [0].PutPos (1, x.num)
  }
  a := x.pg [x.dp].Get (x.idx [x.dp])
  if x.pg [x.dp].GetPos (0) == 0 { // we are on leaf level
tst (1)
    x.pg [x.dp].Del (x.idx [x.dp])
/**********************************************************
//  with pg [dp]
    if idx [dp] + 1 < num {
      for i := idx [dp} + 1; i < num; i++ {
        content [i - 1] = content [i]
        pos [i] = pos [i + 1]
      }
    }
    content [num - 1] = empty
    pos [num] = 0
}
**********************************************************/
    for i := uint(0); i < x.pg [x.dp].GetNum(); i++ {
errh.Error2 ("i =", i, ">>> Pos =", x.pg [x.dp].GetPos (i))
    }
    x.pg [x.dp].DecNum()
    if x.dp == 1 { // page underflow is allowed in the root
      x.write (x.pg [x.dp], 1)
      x.write (x.pg [0], 0)
tst (4)
      return a
    } else { // x.dp > 1
      if x.pg [x.dp].GetNum() < N {
        x.removeUnderflow (x.dp)
      } else {
        x.write (x.pg [x.dp], x.pg [x.dp - 1].GetPos (x.idx [x.dp - 1]))
      }
    }
  } else { // pg [x.dp].pos [0]] > 0, i.e. we are not on leaf level
// We look for the greatest object < x.Get (idx [x.dp]) in a node of depth x.dp,
// copy it to x.content [idx [x.dp]] and replace it by x.empty:
tst (10)
    d := x.dp
    for {
      n := x.idx [d]
      x.pg [d + 1] = x.read (x.pg [d].GetPos (n))
      d++
      if x.pg [d].GetPos (x.idx [d]) == 0 { // we are on leaf level
        x.pg [d].DecNum()
        x.idx [d] = x.pg [d].GetNum()
        if x.idx [d] > 100 { ker.Oops() }
        x.pg [x.dp].Put (x.idx [x.dp], x.pg [d].Get (x.idx [d]))
        x.pg [d].Put (x.idx [d], x.empty)
        break
      }
    }
    x.write (x.pg [d], x.pg [d - 1].GetPos (x.idx [d - 1]))
    x.write (x.pg [x.dp], x.pg [x.dp - 1].GetPos (x.idx [x.dp - 1]))
    if x.pg [d].GetNum() < N {
      x.removeUnderflow (d)
    }
  }
  x.write (x.pg [0], 0)
  if x.num == 0 {
    // TODO
  } else if ! x.ExGeq (a) {
    x.Jump (true)
  }
  return a
}

func (x *persistentSet) Ex (a Any) bool {
  x.dp = 0
  n := x.pg [x.dp].GetPos (x.idx [0])
  for {
    x.dp++
    x.pg [x.dp] = x.read (n)
    if x.pg [x.dp].GetNum() == 0 {
      return false
    }
    x.idx [x.dp] = 0
    i1 := x.pg [x.dp].GetNum()
    for x.idx [x.dp] < i1 {
      i := (x.idx [x.dp] + i1) / 2
      if Less (x.pg [x.dp].Get (i), a) {
        x.idx [x.dp] = i + 1
      } else {
        i1 = i
      }
    }
    if x.idx [x.dp] < x.pg [x.dp].GetNum() {
      if ! Less (a, x.pg [x.dp].Get (x.idx [x.dp])) {
        return true
      }
    }
    n = x.pg [x.dp].GetPos (x.idx [x.dp])
    if n == 0 {
      return false
    }
  }
  return false
}

func (x *persistentSet) Write() {
  const n0 = 1
  for n := uint(0); n < x.seq.Num(); n++ {
    scr.Colours (col.Yellow(), col.Green()); scr.WriteNat (n, n0 + n, 0)
//    x.tmp = x.read (n)
//    x.tmp.Write (n0 + n, 8)
    x.read (n).Write (n0 + n, 4)
  }
  for d := uint(1); d < maxDepth; d++ {
    scr.WriteNat (x.idx [d], n0 + d, 70)
  }
}

func (x *persistentSet) ExGeq (a Object) bool {
  x.dp = 0
  n := x.pg [x.dp].GetPos (0)
  loop:
  for {
    if x.dp >= maxDepth { ker.Oops() }
    x.dp++
    x.pg [x.dp] = x.read (n)
    if x.pg [x.dp].GetNum() == 0 {
      return false
    }
    x.idx [x.dp] = 0
    i1 := x.pg [x.dp].GetNum()
    for x.idx [x.dp] < i1 {
      i := (x.idx [x.dp] + i1) / 2
      if Less (x.pg [x.dp].Get (i), a) {
        x.idx [x.dp] = i + 1
      } else {
        i1 = i
      }
    }
    if x.idx [x.dp] < x.pg [x.dp].GetNum() {
      if Less (a, x.pg [x.dp].Get (x.idx [x.dp])) {
        if x.pg [x.dp].GetPos (x.idx [x.dp]) == 0 {
          return true // found next object
        } else {
          // look deeper
        }
      } else {
        return true // a.Eq (x.pg [x.dp].Get (x.idx [x.dp])
      }
    }
    n = x.pg [x.dp].GetPos (x.idx [x.dp])
    if n == 0 {
      for {
        if x.idx [x.dp] == x.pg [x.dp].GetNum() {
          if x.dp == 1 {
            return false
          } else {
            x.dp--
          }
        } else {
          break loop
        }
      }
    }
  }
  return true
}

func (x *persistentSet) NumPred (p Pred) uint {
  return 0
}

func (x *persistentSet) All (p Pred) bool {
  return false
}

func (x *persistentSet) ExPred (p Pred, f bool) bool {
  return false
}

func (x *persistentSet) StepPred (p Pred, f bool) bool {
  return false
}

func (x *persistentSet) trav (d, n uint, op Op) {
  for i := uint(0); i <= x.pg [d].GetNum(); i++ {
    if x.pg [d].GetPos (i) > 0 {
      x.pg [d + 1] = x.read (x.pg [d].GetPos (i))
      x.trav (d + 1, (2 * N + 1) * n + i, op)
    }
    if i < x.pg [d].GetNum() {
      x.pg[d].Oper (i, op)
    }
  }
}

func (x *persistentSet) Trav (op Op) {
  if x.num == 0 { return }
  x.trav (1, 0, op)
}

func (x *persistentSet) Filter (y Iterator, p Pred) {

}

func (x *persistentSet) Cut (y Iterator, p Pred) {

}

func (x *persistentSet) ClrPred (p Pred) {

}

func (x *persistentSet) Split (y Iterator) {

}

func (x *persistentSet) Join (y Iterator) {

}
