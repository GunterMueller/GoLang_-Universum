package main

import ("murus/scr"; "murus/errh"; "murus/img")

func main() {
  scr.NewMax()
  const a = 100
  var p = []string { "Anaxagoras", "Aristoteles", "Cicero", "Nemo", "Demokrit", "Diogenes",
                     "Epikur", "Heraklit", "Platon", "Protagoras", "Pythagoras", "Sokrates", "Thales" }
  var x, y uint
  for _, s:= range p {
    img.Get (s, x, y)
    errh.Error (s, 0)
    if x + 2 * a < scr.Wd() {
      x += a
    } else {
      x = 0; y = a
    }
  }
  errh.Error2 ("", uint(len(p)), "Philosophen", 0)
  scr.Fin()
}
