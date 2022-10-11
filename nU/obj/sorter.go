package obj

// (c) Christian Maurer   v. 220702 - license see nU.go

type
  Sorter interface {

//
  Ordered () bool

//
  Sort ()

//
  ExGeq (a any) bool
}
