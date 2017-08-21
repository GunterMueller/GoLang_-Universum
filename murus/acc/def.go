package acc

type Account interface {

// Vor.: a > 0
// Der Kontobestand ist um a erhöht.
  Deposit (a uint)

// Vor.: a > 0
// Liefert genau a, wenn der Kontobestand >= a war.
// In diesem Fall ist der Kontobestand um a erniedrigt.
// Liefert andernfalls 0.
  Draw (a uint) uint
}

func New() Account { return new_() }
