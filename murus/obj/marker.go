package obj

// (c) murus.org  v. 140102 - license see murus.go

type
  Marker interface {

// x is marked, iff m.
  Mark (m bool)

// Returns true, iff x is marked.
  Marked () bool
}
