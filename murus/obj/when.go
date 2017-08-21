package obj

// (c) murus.org  v. 170411 - license see murus.go

// guarded selective waiting
func When (b bool, c chan Any) chan Any {
  if b {
    return c
  }
  return nil
}
