package enum

// (c) murus.org  v. 170419 - license see murus.go

var (
  lSparsCode, sSparsCode []string =
  []string { "AAA", "AAD", "ADD", "DAD", "DDD" },
  lSparsCode
  NSparsCodes = uint8(len(lSparsCode))
)

func init() {
  l[SparsCode], s[SparsCode] = lSparsCode, sSparsCode
}
