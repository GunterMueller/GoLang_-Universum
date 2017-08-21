package obj

// (c) murus.org  v. 151212 - license see murus.go

type (
// statements
  Stmt func()
  StmtSpectrum func (uint)
// conditions
  Cond func() bool
  CondSpectrum func (uint) bool
// operations
  Op func (Any)
  OpSpectrum func (Any, uint)
  Op3 func (Any, Any, Any)
// functions
  Func func (Any) Any
  FuncSpectrum func (Any, uint) Any
// predicates
  Pred func (Any) bool
  PredSpectrum func (Any, uint) bool
// conditioned operations
  CondOp func (Any, bool)
  CondOp3 func (Any, Any, Any, bool)
  CondOp3bool func (Any, Any, bool, Any, bool)
// relations
  Rel func (Any, Any) bool
// writings
  Writing func (Any, uint, uint)
  Writing2 func (Any, uint, uint, uint, uint)
)

// Stmt
func Null() { }

// Op
func Ignore (a Any) { }

// Cond[Spectrum]
func True() bool { return true }
func TrueSp (i uint) bool { return true }

// Func[Spectrum]
func Id (a Any) Any { return a }
func IdSp (a Any, i uint) Any { return a }
func Nil (a Any) Any { return nil }
func NilSp (a Any, i uint) Any { return nil }

// Pred[Spectrum]
func AllTrue (a Any) bool { return true }
func AllTrueSp (a Any, i uint) bool { return true }

// CondOp
func CondIgnore (a Any, b bool) { }
func CondIgnore3 (a, a1, a2 Any, b bool) { }
