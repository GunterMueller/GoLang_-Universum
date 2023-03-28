package bahnhof

// (c) Christian Maurer   v. 230107 - license see µU.go

type
  Bahnhof interface {

  Betreiben()
}

func New() Bahnhof { return new_() }
