package acc

// (c) Christian Maurer   v. 171020 - license see µU.go

// XXX funzt nicht

type (
  pair struct {
              uint
            c chan uint
              }
  accountCh struct {
           balance uint
   cDraw, cDeposit chan pair
                   }
)

func newCh() Account {
  var ps []pair
  x := new(accountCh)
  go func() {
    ps = make([]pair, 0)
    for {
      select {
      case p := <-x.cDeposit:
        x.balance += p.uint
        p.c <- x.balance
        for _, p := range ps {     // vorgemerkte Aufträge auszahlen,
          if p.uint <= x.balance { // wenn der Kontobestand das erlaubt
            x.balance -= p.uint
            p.c <- x.balance
          }
        }
      case p := <-x.cDraw:
        if p.uint <= x.balance {
          x.balance -= p.uint
          p.c <- x.balance
        } else {
          ps = append(ps, p) // Auftrag vormerken
        }
      }
    }
  }()
  return x
}

func (x *accountCh) Draw (a uint) uint {
  if a == 0 {
    return x.balance
  }
  p := pair { a, make(chan uint) }
  x.cDraw <- p
  return <-p.c
}

func (x *accountCh) Deposit (a uint) uint {
  if a == 0 {
    return x.balance
  }
  p := pair { a, make(chan uint) }
  x.cDeposit <- p
  return <-p.c
}
