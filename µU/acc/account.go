package acc

type (
  pair struct {
              uint
            c chan uint
              }
  account struct {
           cDraw,
        cDeposit chan pair
                 }
)
var
  busyWaiting = false // true, wenn Kunde "pollen" muss, und
                      // false, wenn Aufträge vorgemerkt werden,

func new_() Account {
  var ps []pair
  x := new(account)
  go func() {
    balance := uint(50) // Lockangebot
    if ! busyWaiting {
      ps = make([]pair, 0)
    }
    for {
      select {
      case p := <-x.cDraw:
        if p.uint <= balance {
          balance -= p.uint
          p.c <- p.uint
        } else {
          if busyWaiting {
            p.c <- 0 // Auszahlung nicht möglich;
                     // Kunde muss es später wieder versuchen
          } else {
            ps = append(ps, p) // Auftrag vormerken
          }
        }
      case p := <-x.cDeposit:
        balance += p.uint
        if ! busyWaiting {
          for _, p := range ps {   // vorgemerkte Aufträge auszahlen,
            if p.uint <= balance { // wenn der Kontobestand das erlaubt
              balance -= p.uint
              p.c <- p.uint
            }
          }
        }
      }
    }
  }()
  return x
}

func (x *account) Draw (a uint) uint {
  if a == 0 { return 0 }
  p := pair { a, make (chan uint) }
  x.cDraw <- p
  return <-p.c
}

func (x *account) Deposit (a uint) {
  if a == 0 { return }
  x.cDeposit <- pair { a, nil }
}
