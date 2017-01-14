package lockp

// (c) murus.org  v. 161212 - license see murus.go

// >>> Tournament-Algorithm of Kessels due to Taubenfeld, p. 41

type
  kesselsN struct {
       nProcesses uint
       interested [][2]bool
             turn [][2]uint
                b []uint
                  }

func newKN (n uint) LockerP {
  if n < 2 { return nil }
  x := new (kesselsN)
  x.nProcesses = n
  x.b = make ([]uint, n)
  x.interested = make ([][2]bool, n)
  x.turn = make ([][2]uint, n)
  return x
}

func (x *kesselsN) Lock (i uint) {
  if i >= x.nProcesses { return }
  n := i + x.nProcesses
  for n > 1 {
    j := n % 2
    n /= 2
    x.interested[n][j] = true
    local := (x.turn[n][1 - j] + j) % 2
    x.turn[n][j] = local
    for x.interested[n][1 - j] && local == (x.turn[n][1 - j] + j) % 2 {
      // Null()
    }
    x.b[n] = j
  }
}

func (x *kesselsN) Unlock (i uint) {
  n := uint(1)
  for n < x.nProcesses {
    x.interested[n][x.b[n]] = false
    n = 2 * n + x.b[n]
  }
}
