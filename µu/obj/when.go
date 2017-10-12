package obj

// (c) Christian Maurer   v. 170411 - license see µu.go

// guarded selective waiting
func When (b bool, c chan Any) chan Any {
  if b {
    return c
  }
  return nil
}
