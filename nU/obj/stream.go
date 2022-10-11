package obj

// (c) Christian Maurer   v. 220702 - license see nU.go

type (
  Stream = []byte
  BoolStream = []bool
  IntStream = []int
  UintStream = []uint
  AnyStream = []any
)

func Streamic (a any) bool {
  if a == nil { return false }
  switch a.(type) {
  case Stream, BoolStream, IntStream, UintStream, AnyStream:
    return true
  }
  return false
}
