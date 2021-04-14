package reg

// (c) Christian Maurer   v. 210331 - license see µUu.go

type
  Register interface { // Register mit ganzzahligen Werten.
                       // Für alle Methoden gilt als Vor., dass das aufrufende
                       // Register durch eine Wertzuweisung von  Null() oder
                       // Wert einer registerwertigen Funktion erzeugt wurde.
                       // Das aufrufende Objekt wird immer mit "x" bezeichnet.

// Vor.: Der Wert von x ist jetzt um 1 erhöht.
  Inc()

// Vor.: x hat einen Wert > 0.
// Sein Wert ist jetzt um 1 erniedrigt.
  Dec()

// Liefert genau dann true, wenn x einen Wert > 0 hat.
  Gt0() bool

// Liefert ein neues Register mit dem Wert
// der Summe der Werte von a und x.
  Add (a Register) Register

// Liefert ein neues Register mit dem Wert
// des Produkts der Werte von a und x.
  Mul (a Register) Register

// Der Wert von x ist in einer eigenen Zeile auf dem Bildschirm ausgegeben.
  Write()
}

// Liefert ein neues Register mit dem Wert 0.
func Null() Register { return null() }

type Registers interface { // Folgen von Registern

// Liefert genau dann true, wenn x nicht leer ist,
// d.h., wenn sie mindestens ein Register enthält.
  NotEmpty() bool

// Liefert ein Register mit dem Wert der Anzahl der Register von x.
  Num() Register

// Vor.: x ist nicht leer.
// Liefert das erste Register von x.
  Head() Register

// Vor.: x ist nicht leer.
// Liefert den Rest von x, d.h.
// x ohne das vorher erste Register von x.
  Tail() Registers

// Liefert die Folge mit r als erstem Register
// und der x als Rest.
  Cons (r Register) Registers
}

// Vor.: x ist entweder leer oder besteht nur aus einem Register oder
//       aus mehreren voneinander durch Kommas getrennten Registern.
// Liefert eine Folge vom Typ Registers.
// Sie ist leer, falls a leer war; ansonsten enthält sie genau
// die übergebenen Register in der angegebenen Reihenfolge.
func New (a ...Register) Registers { return new_(a...) }

type RegFunc func (a Registers) Register
type RegFunc1 func (a Register, as Registers) Register
