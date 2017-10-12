package enum

// (c) Christian Maurer   v. 170419 - license see µu.go

var (
  lSparsCode, sSparsCode []string =
  []string { "AAA", "AAD", "ADD", "DAD", "DDD" },
  lSparsCode
  NSparsCodes = uint8(len(lSparsCode))
)

func init() {
  l[SparsCode], s[SparsCode] = lSparsCode, sSparsCode
}
