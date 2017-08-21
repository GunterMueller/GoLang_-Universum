package obj

// (c) murus.org  v. 170701 - license see murus.go

type
  Object interface {

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

// Returns true, iff the type of a is bool,
// [u]int{8|16|32}, float[32|64], complex[64|128],
// string or []byte (we treat them also as atomic).
func Atomic (a Any) bool { return atomic(a) }

// Returns true, iff the type of a implements Object.
func IsObject (a Any) bool { return isObject(a) }

// Returns true, iff a is atomic or implements Object
// (the types that are particularly supported by murus).
func AtomicOrObject (a Any) bool {
  return atomic (a) || isObject (a)
}
