package dgra

// (c) Christian Maurer   v. 200717 - license see µU.go

import
  "µU/errh"
const
  withError = false

func (x *distributedGraph) log0 (s string) {
  if ! x.demo { return }
  if withError {
    errh.Error0 (s)
    return
  }
  println(s)
}

func (x *distributedGraph) log (s string, n uint) {
  if ! x.demo { return }
  if withError {
    errh.Error (s, n)
    return
  }
  println(s, n)
}

/*
func (x *distributedGraph) enter (r uint) {
  if ! x.demo { return }
  s := "after round"
  if withError {
    errh.Error (s, r)
    return
  }
  println (s, r)
}
*/

func (x *distributedGraph) log2 (s string, n uint, s1 string, n1 uint) {
  if ! x.demo { return }
  if withError {
    errh.Error2 (s, n, s1, n1)
    return
  }
  println(s, n, s1, n1)
}

func (x *distributedGraph) log3 (s string, n uint, s1 string, n1 uint, s2 string, n2 uint) {
  if ! x.demo { return }
  if withError {
    errh.Error3 (s, n, s1, n1, s2, n2)
    return
  }
  println(s, n, s1, n1, s2, n2)
}

func (x *distributedGraph) log4 (s string, n uint, s1 string, n1 uint, s2 string, n2 uint, s3 string, n3 uint) {
  if ! x.demo { return }
  if withError {
    errh.Error4 (s, n, s1, n1, s2, n2, s3, n3)
    return
  }
  println(s, n, s1, n1, s2, n2, s3, n3)
}
