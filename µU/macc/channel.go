package macc

// (c) Christian Maurer   v. 190823 - license see µU.go

// XXX funzt nicht

type (
  pair struct {
              uint
            c chan uint
              }
  maccountCh struct {
            balance uint
    cDraw, cDeposit chan pair
                    }
)

func newChannel() MAccount {
  var ps []pair
  x := new(maccountCh)
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

func (x *maccountCh) Draw (a uint) uint {
  if a == 0 {
    return x.balance
  }
  p := pair { a, make(chan uint) }
  x.cDraw <- p
  return <-p.c
}

func (x *maccountCh) Deposit (a uint) uint {
  if a == 0 {
    return x.balance
  }
  p := pair { a, make(chan uint) }
  x.cDeposit <- p
  return <-p.c
}
