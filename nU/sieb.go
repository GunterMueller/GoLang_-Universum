package main // Sieb des Eratosthenes

const N = 313

func start (out chan uint) {
  a := uint(2)
  out <- a
  a++
  for {
    out <- a
    a += 2
  }
}

func sieve (in, out, off chan uint) {
  Primzahl := <-in
  off <- Primzahl
  for {
    a := <-in
    if (a % Primzahl) != 0 { out <- a }
  }
}

func stop (in chan uint) {
  for { <-in }
}

func write (in chan uint, d chan bool) {
  for i := 1; i < N; i++ {
    print(<-in, " ")
  }
  println()
  d <- true
}

func main() {
  var c [N]chan uint
  for i := 0; i < N; i++ {
    c[i] = make(chan uint)
  }
  out := make(chan uint)
  done := make(chan bool)
  go start (c[0]) // Start
  for i := 1; i < N; i++ {
    go sieve (c[i-1], c[i], out) // Sieb[i]
  }
  go stop (c[N - 1]) // Ende
  go write (out, done) // Ausgabe
  <-done
}
