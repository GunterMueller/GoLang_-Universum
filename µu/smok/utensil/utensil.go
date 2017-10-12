package utensil

// (c) Christian Maurer   v. 150818 - license see µu.go

import
  "µu/col"
var (
  text   = [NUtensils]string     {"Papier",        " Tabak",        "Hölzer" }
  colour = [NUtensils]col.Colour {col.LightWhite(), col.LightBrown(), col.Sandgelb1()}
)

func other1 (u uint) uint {
  return (u + 1) % NUtensils
}

func other2 (u uint) uint {
  return (u + NUtensils - 1) % NUtensils
}

func others (u uint) (uint, uint) {
  return (u + 1) % NUtensils, (u + NUtensils - 1) % NUtensils
}
