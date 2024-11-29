package main

// (c) Christian Maurer   v. 241019 - license see µU.go

import (
  "os"
  "µU/ego"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/host"
  "µU/euro"
  "µU/macc"
)
var
  acc macc.MAccount

func mask() {
  scr.Colours (col.FlashWhite(), col.Black())
  scr.Write (" balance            Euro ", 1, 0)
  scr.Write ("  amount            Euro ", 2, 0)
  errh.Hint (" deposit: +      draw: - ")
}

func main() {
  me := ego.Me()
  w, h := uint(200), uint(80)
  const (lb = 1; la = 2; c = 9)
  x := me * (w + 4)
  scr.NewWH (x, 0, w, h); defer scr.Fin()
  acc = macc.New (host.LocalName(), 50888, me == 0)
  balance, amount := euro.New(), euro.New()
  d, v := false, uint(0)
  if me == 0 {
    errh.Hint ("I am the server")
    for {
      if cmd, _ := kbd.Command(); cmd == kbd.Esc {
        os.Exit(0)
      }
    }
  }
  mask()
  for {
    v = acc.Show (0)
    balance = euro.New2 (v / 100, v % 100)
    balance.Write (lb, c)
    loop:
    for {
      b, cmd, _ := kbd.Read()
      if cmd == kbd.Esc {
        os.Exit (0)
      }
      switch b {
      case '+':
        d = true // deposit
        break loop
      case '-':
        d = false // draw
        break loop
      }
    }
    v = acc.Show (0)
    balance = euro.New2 (v / 100, v % 100)
    balance.Write (lb, c)
    amount.Edit (la, c)
    if d {
      v = acc.Deposit (amount.Val()) // amount = 100
      balance = euro.New2 (v / 100, v % 100)
      balance.Write (lb, c)
      amount.Clr()
      amount.Write (la, c)
    } else {
      v = acc.Draw (amount.Val())
      if v == 0 {
        errh.Error0 ("balance is insufficient")
      } else {
        balance = euro.New2 (v / 100, v % 100)
        balance.Write (lb, c)
        amount.Clr()
        amount.Write (la, c)
      }
    }
  }
}
