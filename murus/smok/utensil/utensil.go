package utensil

// (c) murus.org  v. 150513 - license see murus.go

import
  "murus/col"
var (
  text   = [NUtensils]string    { "Papier",       " Tabak",       "Hölzer" }
  colour = [NUtensils]col.Colour{ col.LightWhite, col.LightBrown, col.Sandgelb1}
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
