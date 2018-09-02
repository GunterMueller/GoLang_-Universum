package obj

// (c) Christian Maurer   v. 180902 - license see nU.go

type (
  Stream = []byte
  BoolStream = []bool
  IntStream = []int
  UintStream = []uint
  AnyStream = []Any
)

func Streamic (a Any) bool {
  if a == nil { return false }
  switch a.(type) {
  case Stream, BoolStream, IntStream, UintStream, AnyStream:
    return true
  }
  return false
}
