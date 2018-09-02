package obj

// (c) Christian Maurer   v. 180902 - license see nU.go

type Object interface {

// Most objects in computer science can be compared with others,
// whether they are equal, and can be copied, so they have the type
  Equaler // see equaler.go

// Furthermore, usually we can order objects; so they have the type
  Comparer // see comparer.go

// Moreover they can be empty and may be cleared with the effect
// of being empty, hence they have the type
  Clearer // see clearer.go

// and can be serialized into connected byte sequences,
// e.g. to be written to a storage device or transmitted
// over communication channels, so they have the type
  Coder // see coder.go
}

// Returns true, iff the type of a implements Object.
func IsObject (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Object)
  return ok
}

// Returns true, iff the type of a is bool,
// [u]int{8|16|32}, float[32|64], complex[64|128] or string.
func Atomic (a Any) bool {
  switch a.(type) {
  case bool,
       int8,  int16,  int32,  int,  int64,
       uint8, uint16, uint32, uint, uint64,
       float32, float64, complex64, complex128,
       string:
    return true
  }
  return false
}

// Returns true, iff the type of a is atomic or implements Object
// (the types that are particularly supported by nU).
func AtomicOrObject (a Any) bool {
  return Atomic (a) || IsObject (a)
}
