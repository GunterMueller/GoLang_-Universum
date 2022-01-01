package ker

// (c) Christian Maurer   v. 211218 - license see µU.go

func binom (n, k uint) uint {
  if k == 1 { return n }
  if n < k { return 0 }
  if n == k { return 1 }
  if n - k < k { k = n - k }
  b := uint(1)
  for i := uint(1); i <= k; i++ {
    b *= n
    b /= i
    n --
  }
  return b
}
