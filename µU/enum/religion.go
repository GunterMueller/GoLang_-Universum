package enum

// (c) Christian Maurer   v. 201007 - license see µU.go

var
  lReligion = []string {"keine",
                        "evangelisch",
                        "katholisch",
                        "jüdisch",
                        "muslimisch",
                        "hinduistisch",
                        "buddhistisch",
                        "andere",
                       }

func init() {
  l[Religion] = lReligion
  s[Religion] = lReligion
}
