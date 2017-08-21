package enum

// (c) murus.org  v. 140522 - license see murus.go

var (
  lReligion, sReligion =
  []string { "keine", "evangelisch", "katholisch", "jüdisch", "muslimisch",
             "hinduistisch", "buddhistisch", "andere" },
  lReligion
  NReligions = uint8(len(lReligion))
)

func init() {
  l[Religion], s[Religion] = lReligion, sReligion
}
