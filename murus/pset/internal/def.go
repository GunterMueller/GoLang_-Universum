package internal

// (c) Christian Maurer   v. 170316 - license see murus.go

import
  . "murus/obj"
const (
  N = 2 // provisorial
)
type
  Page interface {

  Object

  PutNum (n uint)
  GetNum () uint
  PutPos (p, n uint)
  GetPos (p uint) uint
  Put (p uint, o Object)
  Get (p uint) Object
  Oper (p uint, op Op)
  Ins (o Object, p, n uint)
  IncNum ()
  DecNum ()
  RotLeft ()
  RotRight ()
  Join (p uint)
  Del (p uint)
  ClrLast ()
  Write (l, c uint)
}

func New(a Any) Page { return new_(a) }
