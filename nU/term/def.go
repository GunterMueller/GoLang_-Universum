package term

// (c) Christian Maurer   v. 171229 - license see nU.go

const (
  Enter = 13; Esc = 27; Tab = 9; Back = 127
)

type Terminal interface {

  Read() byte

  Fin()
}

func New() Terminal { return new_() }
