package show

// (c) Christian Maurer   v. 191019 - license see µU.go

type
  Mode int; const (
  Look = iota
  Walk
  Fly
)

func Arg() Mode { return arg() }
