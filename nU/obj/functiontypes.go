package obj

// (c) Christian Maurer   v. 220702 - license see nU.go

import
  "runtime"
type (
// statements
  Stmt func()
  StmtSpectrum func (uint)
// conditions
  Cond func() bool
  CondSpectrum func (uint) bool
// Nvalued functions
  NFunc func() uint
  NFuncSpectrum func (uint) uint
// operations
  Op func (any)
  OpSpectrum func (any, uint)
// functions
  Func func (any) any
  FuncSpectrum func (any, uint) any
// predicates
  Pred func (any) bool
  PredSpectrum func (any, uint) bool
// conditioned operations
  CondOp func (any, bool)
  CondOp2 func (any, any, bool)
  CondOpSpectrum func (any, bool, uint)
// relations
  Rel func (any, any) bool
  RelSpectrum func (any, any, uint) bool
)

// Stmt[Spectrum]
func Nothing() { runtime.Gosched() }
func NothingSp (i uint) { }

// Op[Spectrum]
func Ignore (a any) { }
func IgnoreSp (a any, i uint) { }

// NFunc[Spectrum]
func Null() uint { return 0 }
func NullSp(i uint) uint { return 0 }

// Cond[Spectrum]
func True() bool { return true }
func TrueSp (i uint) bool { return true }

// Func[Spectrum]
func Id (a any) any { return a }
func IdSp (a any, i uint) any { return a }
func Nil (a any) any { return nil }
func NilSp (a any, i uint) any { return nil }

// Pred[Spectrum]
func AllTrue (a any) bool { return true }
func AllTrueSp (a any, i uint) bool { return true }

// CondOp[Spectrum]
func CondIgnore (a any, b bool) { }
func CondIgnore2 (a, a1 any, b bool) { }
func CondIgnoreSp (a any, b bool, i uint) { }

// we get rid of TravPred by:
func PredOp2Op (p Pred, o Op) Op {
  return func (a any) { if p(a) { o(a) } }
}

// we get rid of TravCond by:
func PredCondOp2Op (p Pred, o CondOp) Op {
  return func (a any) { if p(a) { o(a, true) } else { o(a, false) } }
}
