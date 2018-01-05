package internal
import . "nU/obj"

const (Candidate = byte(iota); Reply; Leader)

type Message interface {

  Equaler
  Coder

// Liefert den Typ von x.
  Type() byte

// Liefert das Quadrupel (id, num, maxnum, ok) von x.
  IdNumsOk() (uint, uint, uint, bool)

// x besteht aus Typ Candidate, id i, num n, maxnum m
// und undefiniertem ok.
  SetPass (i, n, m uint)

// x besteht aus Typ Reply und ok b,
// die anderen Komponenten sind unverändert.
  SetReply (b bool)

// x besteht aus Typ Leader and id i,
// die anderen Komponenten sind unverändert.
  SetLeader (i uint)
}

// Liefert eine neue Botschaft, bestehend
// aus zero values in allen Komponenten.
func New() Message { return new_() }
