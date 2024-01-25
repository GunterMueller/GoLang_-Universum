package schan

// (c) Christian Maurer   v. 220801 - license see nU.go

type
  SynchronousChannel interface {

  Send (a any)
  Recv() any
}

func New (a any) SynchronousChannel { return new_(a) }
