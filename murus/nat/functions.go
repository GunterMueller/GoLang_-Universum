package nat

// (c) Christian Maurer  v. 121119 - license see murus.go

func sumDigits (n uint) uint {
  s:= uint(0)
  for n > 0 {
    s += n % 10
    n /= 10
  }
  return s
}

func gcd (n, k uint) uint {
  if n < k { n, k = k, n }
  if k == 0 { return n }
  return gcd (n % k, k)
}

func lcm (n, k uint) uint64 {
  if n == 0 || k == 0 { return 0
  }
  return (uint64(n) * uint64 (k)) / uint64(gcd (n, k))
}

func fak (n uint) uint {
  a:= uint(1)
  for i:= uint(2); i <= n; i++ {
    a *= i
  }
  return a
}

func p (n, k, a uint) uint {
  if k == 0 { return a }
  if k % 2 == 0 {
    return p (n * n, k / 2, a)
  }
  return p (n * n, k / 2, n * a)
}

func pow (n, k uint) uint {
  return p (n, k, 1)
}

func lowFak (n, k uint) uint {
  if n < k { return 0 }
  a:= uint(1)
  for i:= uint(n - k + 1); i <= n; i++ {
    a *= i
  }
  return a
}

func stirl2 (n, k uint) uint {
  a, b:= 0, 1
  e:= k % 2 == 1
  for i:= uint(1); i <= k; i++ {
    b *= int(k - i + 1)
    b /= int(i)
    if e {
      a += b * int(pow (i, n))
    } else {
      a -= b * int(pow (i, n))
    }
    e = ! e
  }
  return uint(a / int(fak (k)))
}
