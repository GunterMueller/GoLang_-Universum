package dgra

// (c) Christian Maurer   v. 170504 - license see murus.go

import
  "murus/errh"
var
  ohneError bool

func init() {
  ohneError = true
  ohneError = false
}

func (x *distributedGraph) log0 (s string) {
  if ! x.demo { return }
  println(s)
  if ohneError { return }
  errh.Error0(s)
}

func (x *distributedGraph) log (s string, n uint) {
  if ! x.demo { return }
  println(s, n)
  if ohneError { return }
  errh.Error(s, n)
}

func (x *distributedGraph) log2 (s string, n uint, s1 string, n1 uint) {
  if ! x.demo { return }
  println(s, n, s1, n1)
  if ohneError { return }
  errh.Error2(s, n, s1, n1)
}

func (x *distributedGraph) log3 (s string, n uint, s1 string, n1 uint, s2 string, n2 uint) {
  if ! x.demo { return }
  println(s, n, s1, n1, s2, n2)
  if ohneError { return }
  errh.Error3(s, n, s1, n1, s2, n2)
}

func (x *distributedGraph) enter (r uint) {
  if ! x.demo { return }
  println("start round", r)
  if ohneError { return }
  errh.Error ("start round", r)
}

func (x *distributedGraph) end (r uint) {
  if ! x.demo { return }
  println("round", r, "ended")
  if ohneError { return }
  errh.Error("end of round", r)
}
