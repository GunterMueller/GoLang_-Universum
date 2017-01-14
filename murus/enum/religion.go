package enum

// (c) murus.org  v. 140522 - license see murus.go

var
  lReligion, sReligion =
  []string { "keine", "evangelisch", "katholisch", "jüdisch", "muslimisch",
             "hinduistisch", "buddhistisch", "andere" },
  lReligion
const
  NReligions = 8


func init() {
  l[Religion], s[Religion] = lReligion, sReligion
}
