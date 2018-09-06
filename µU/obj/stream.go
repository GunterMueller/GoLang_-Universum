package obj

// (c) Christian Maurer   v. 180812 - license see µU.go

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
  case BoolStream, Stream, IntStream, UintStream, AnyStream:
    return true
  }
  return false
}
