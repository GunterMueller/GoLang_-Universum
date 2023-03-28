package kilo

// (c) Christian Maurer   v. 230106 - license see µU.go

type
  Kilometrierung byte; const (
  Mit = Kilometrierung(iota) // von links nach rechts
  Gegen
  NK
)
var
  Ktext = [NK+1]string {"Mit", "Gegen", "NK"}

// Liefert Gegen für K == Mit, sonst Mit.
func Gegenrichtung (k Kilometrierung) Kilometrierung { return entgegen(k) }
