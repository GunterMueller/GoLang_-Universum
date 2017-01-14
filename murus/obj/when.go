package obj

// (c) murus.org  v. 140102 - license see murus.go

// selective communication
func When (b bool, c chan Any) chan Any {
  if b {
    return c
  }
  return nil
}
