package obj

// (c) Christian Maurer   v. 220420 - license see µU.go

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
  case BoolStream, Stream, IntStream, UintStream, AnyStream:
    return true
  }
  return false
}
