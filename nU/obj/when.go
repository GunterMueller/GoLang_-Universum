package obj

// (c) Christian Maurer   v. 220702 - license see nU.go

// guarded selective waiting
func When (b bool, c chan any) chan any {
  if b {
    return c
  }
  return nil
}
